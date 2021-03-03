package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/spf13/pflag"
	tmpl "github.com/ukautz/tmpl/pkg"
	_ "github.com/ukautz/tmpl/pkg/imports"
)

var Version = "0.1.0"
var BuildCommit = "UNSET"

const USAGE_HEAD = `tmpl renders documents from templates and structured data

 version     %s
 commit      %s
 source      https://github.com/ukautz/tmpl

Arguments

`

const USAGE_FOOT = `

Data Locations:
---------------
Multiple data locations are accepted:
- Local files:
  - file://<path-to-file> like: "file:///etc/config.json"
  - /path/to/file, like "/etc/config.json" (must exist)
- HTTP sources:
  - like "https://domain.tld/file.json"
- Environment variables:
  - without prefix, like "env:", will create data out of all environment variables
  - with prefix, like "env:APP_", will create data only of environment variables with given prefix
- Command execution:
  - shell://<arbitrary-shell-command> like: "shell://ps auxf"

Decoders:
---------

Template:
---------

Rendered:
---------

Output:
-------

Documentation:
--------------
  https://github.com/ukautz/tmpl/blob/%s/README.md
`

func fail(msg string) {
	fmt.Printf("\033[1;31mFAILED: %s\033[0m\n\n", msg)
	usage(1)
}

func usage(exit int) {
	fmt.Printf(USAGE_HEAD, Version, BuildCommit)
	args := []string{}
	for arg, val := range map[string]string{
		"data":     "location",
		"template": "location",
		"renderer": "name",
		"decoder":  "name",
		"output":   "path",
	} {
		a := pflag.Lookup(arg)
		p := fmt.Sprintf("--%s | -%s <%s>", arg, a.Shorthand, val)
		s := strings.Repeat(" ", 30-len(p))
		args = append(args, fmt.Sprintf("  %s %s %s", p, s, a.Usage))
	}
	sort.Strings(args)
	fmt.Println(strings.Join(args, "\n"))
	fmt.Printf("  %-30s  Show this help\n", "--help | -h")
	fmt.Printf(USAGE_FOOT, Version)
	os.Exit(exit)
}

func main() {
	data := pflag.StringP("data-location", "d", "", "Location of the data to be used for rendering template")
	decoder := pflag.StringP("decoder", "D", "guess", "Which decoder to use for data (default: guess from data location)")
	template := pflag.StringP("template-location", "t", "", "Location of template that is to be rendered with data")
	renderer := pflag.StringP("renderer", "r", "guess", "Set template render engine like template or pongo2 (default: guess from template location)")
	output := pflag.StringP("output", "o", "-", "Path to output or \"-\" for STDOUT (default: -)")
	help := pflag.BoolP("help", "h", false, "Show help")
	pflag.Parse()

	if *help {
		usage(0)
	}
	if *data == "" {
		fail("Missing --data location")
	}
	if *decoder == "" {
		fail("Missing --decoder name")
	}
	if *template == "" {
		fail("Missing --template location")
	}
	if *renderer == "" {
		fail("Missing --renderer name")
	}
	if *output == "" {
		fail("Missing --output location")
	}

	p, err := tmpl.Build(*data, *decoder, *template, *renderer)
	if err != nil {
		fail(err.Error())
	}

	var w io.Writer
	if *output == "-" {
		w = os.Stdout
	} else if fw, err := os.OpenFile(*output, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644); err != nil {
		fail(fmt.Sprintf("could not open output file \"%s\" for write: %s", *output, err))
	} else {
		defer fw.Close()
		w = fw
	}

	if res, err := p.Produce(); err != nil {
		fail(err.Error())
	} else if _, err = w.Write(res); err != nil {
		fail(fmt.Sprintf("could not write to output file \"%s\": %s", *output, err))
	}
}

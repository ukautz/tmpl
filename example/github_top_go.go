package main

import (
	"github.com/ukautz/tmpl"
	"os"
	"fmt"
	_ "github.com/ukautz/tmpl/imports"
)

const (
	TOP_GO_REPOS = "https://api.github.com/search/repositories?q=language:go&sort=starts&order=desc"
	TEMPLATE = "https://raw.githubusercontent.com/ukautz/tmpl/master/examples/github_top_go.pongo"
)

func main() {
	template := os.Getenv("TEMPLATE")
	if template == "" {
		template = TEMPLATE
	}
	q, err := tmpl.BuildTmpl(TOP_GO_REPOS, "json", template, "pongo2")
	if err != nil {
		panic(err)
	} else if res, err := q.Produce(); err != nil {
		panic(err)
	} else {
		fmt.Println(string(res))
	}
}

// +build example
package main

import (
	_ "embed"
	"fmt"

	tmpl "github.com/ukautz/tmpl/pkg"
	_ "github.com/ukautz/tmpl/pkg/imports"
)

const (
	TOP_GO_REPOS = "https://api.github.com/search/repositories?q=language:go&sort=starts&order=desc"
	TEMPLATE     = "https://raw.githubusercontent.com/ukautz/tmpl/master/examples/github_top_go.pongo"
)

//go:embed github_top_go.pongo
var template string

func main() {
	q, err := tmpl.Build(TOP_GO_REPOS, "json", fmt.Sprintf("template:%s", template), "pongo2")
	if err != nil {
		panic(err)
	} else if res, err := q.Produce(); err != nil {
		panic(err)
	} else {
		fmt.Println(string(res))
	}
}

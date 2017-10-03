package main

import (
	"github.com/ukautz/tmpl/json"
	"fmt"
	"github.com/ukautz/tmpl/template"
)

var dataTemlpate = `{
  "name": "www.acme.tld",
  "aliases": ["acme.tld", "blog.acme.tld"],
  "directory": "/var/www/homepage",
  "directories": [
    {"path": "/foo", "users": ["bar", "baz"]},
    {"path": "/lorem"}
  ]
}`

var templateTemplate = `<VirtualHost {{.data.name}}:80>
    ServerName {{.data.name}}
    ServerAlias{{range .data.aliases}} {{.}}{{end}}
    DocumentRoot "{{.data.directory}}"
    {{range $idx, $directory := .data.directories}}
    <Directory "{{$.data.directory}}{{$directory.path}}">
        Require {{if $directory.users}}user{{range $directory.users}} {{.}}{{end}}{{else}}valid-user{{end}}
    </Directory>{{end}}
</VirtualHost>`

func main() {
	data, err := json.NewDecoder().Decode([]byte(dataTemlpate))
	if err != nil {
		panic(err)
	}
	renderer := &template.Renderer{}
	res, err := renderer.Render(data, []byte(templateTemplate))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}

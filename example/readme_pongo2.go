// +build example
package main

import (
	"fmt"

	tmpl "github.com/ukautz/tmpl/pkg"
	"github.com/ukautz/tmpl/pkg/pongo2"
)

var dataPongo2 = `{
  "name": "www.acme.tld",
  "aliases": ["acme.tld", "blog.acme.tld"],
  "directory": "/var/www/homepage",
  "directories": [
    {"path": "/foo", "users": ["bar", "baz"]},
    {"path": "/lorem"}
  ]
}`

var templatePongo = `<VirtualHost {{ data.name }}:80>
    ServerName {{ data.name }}
    ServerAlias {{ data.aliases | join:" " }}
    DocumentRoot "{{ data.directory }}"
    {% for directory in data.directories %}
    <Directory "{{ data.directory }}{{ directory.path }}">
        Require {% if directory.users %}user {{ directory.users | join:" " }}{% else %}valid-user{% endif %}
    </Directory>{% endfor %}
</VirtualHost>`

func main() {
	data, err := tmpl.NewJSONDecoder().Decode([]byte(dataPongo2))
	if err != nil {
		panic(err)
	}
	renderer := &pongo2.Renderer{}
	res, err := renderer.Render(data, []byte(templatePongo))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}

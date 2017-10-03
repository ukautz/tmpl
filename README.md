tmpl
====

tmpl is a lightweight, shell script friendly document generator.

tmpl is inspired by recent extensions of the shell toolset like [jq](https://github.com/stedolan/jq), which allows to
work with complex, structured (JSON) data from the shell. This tool provides a solution for a special case in that
particular niche: automated generation of complex documents out of structured data using smart templates.

```
$ tmpl --data https://config.acme.tld/node/$(uname -n).json \
    --template /etc/apache/templates/homepage.tmpl \
    --output /var/www/home/index.html

$ tmpl --data "shell://elasticdump ... args ..." \
    --decoder json \
    --template "https://wiki.acme.tld/templates/report.tmpl?auth=$(generate-super-token.sh)" \
    --output - | sendmail ...

$ gen-vhost-data.sh | tmpl -d - -D json -t /etc/apache/vhost.tmpl -o - | install-and-restart.sh
```

There are many possible applications or integrations. Here are a few:

* **Reporting**, eg generate HTML reports or summaries based on "native JSON sources" like ElasticSearch or MongoDB
* **Build pipeline**, eg transform generic YAML/JSON into domain specific configuration files (think YAML -> NGINX config)
* **Monitoring / alerting**, eg generate system alert emails with complex information rendered in a human readable way
* **Transactional messaging**, eg create rich documents from raw JSON (eg JSON -> DocX)
* you can think of something ..

Installation
-------------

Download binary from the [releases page](https://github.com/ukautz/tmpl/releases) or:

```
$ go get github.com/ukautz/tmpl
$ glide get github.com/ukautz/tmpl
$ <package-manager> get github.com/ukautz/tmpl
```

Sources
-------

A source is an URL like `http://acme.tld/foo.json` or `file:///etc/foo.json`. The URL provides the location and
indicates a format.

**Supported formats** are JSON and YAML. Per default, tmpl tries to guess the format from the URL. The decoder can be set
explicitly with `--decoder <json|yaml>` (or `-D <json|yaml>`).

**Supported locations** are:

* `http://` or `https://`: arbitrary, GETable HTTP(S) URLs
* `file:///path/to/file` or `/path/to/file`: arbitrary local files
* `shell://`: arbitrary, atomic shell command lines, eg `shell://date +%F`, which would execute `date +%F` and use the
   STDOUT. Don't use pipes or somesuch..
* `-`: STDIN

Templates
---------

tmpl supports two template render engins: [**template**, from the go standard libraries](https://golang.org/pkg/text/template/)
and the [**pongo2** go library](https://github.com/flosch/pongo2) which implements
[Python Django's templating engine](https://docs.djangoproject.com/en/dev/topics/templates/).

The same as with data sources: specify the renderer explicitly with `--renderer | -r <name>` or let tmpl try to guess
it from the template location:

* Template URLs with `.tmpl` or `.template` file names default to `template` engine. Examples:
  * http://acme.tld/templates/vhost.tmpl?foo=bar
  * file:///etc/apache/vhost.template
* Template URLs with `.pongo2` or `.pongo` files names default to `pongo2` engine/ Examples:
  * http://acme.tld/templates/vhost.pongo2?t=123456
  * file:///etc/apache/vhost.pongo

For the template engine examples below, assume the following data structure & content:

```json
{
  "name": "www.acme.tld",
  "aliases": ["acme.tld", "blog.acme.tld"],
  "directory": "/var/www/homepage",
  "directories": [
    {"path": "/foo", "users": ["bar"]},
    {"path": "/lorem"}
  ]
}
```

and the following expected result (+/- a few empty lines, see [optimized templates here](https://github.com/ukautz/tmpl/tree/master/example)):

```$xslt
<VirtualHost www.acme.tld:80>
    ServerName www.acme.tld
    ServerAlias acme.tld blog.acme.tld
    DocumentRoot "/var/www/homepage"

    <Directory "/var/www/homepage/foo">
        Require user foo bar
    </Directory>
    <Directory "/var/www/homepage/lorem">
        Require valid-user
    </Directory>
</VirtualHost>
```

### `template` engine

```
<VirtualHost {{.data.name}}:80>
    ServerName {{.data.name}}
    ServerAlias{{range .data.aliases}} {{.}}{{end}}
    DocumentRoot "{{.data.directory}}"
    {{range $idx, $directory := .data.directories}}
    <Directory "{{$.data.directory}}{{$directory.path}}">
        {{if $directory.users}}
        Require user{{range $directory.users}} {{.}}{{end}}
        {{else}}
        Require valid-user
        {{end}}
    </Directory>
    {{end}}
</VirtualHost>
```

Find more examples for [including additional templates, working with macro like blocks, ..](https://golang.org/pkg/text/template/#hdr-Actions)

### `pongo2` engine

Example template:

```
<VirtualHost {{ data.name }}:80>
    ServerName {{ data.name }}
    ServerAlias {{ data.aliases | join:" " }}
    DocumentRoot "{{ data.directory }}"
    {% for directory in data.directories %}
    <Directory "{{ data.directory }}{{ directory.path }}">
        {% if directory.users %}
        Require user {{ directory.users | join:" " }}
        {% else %}
        Require valid-user
        {% endif %}
    </Directory>
    {% endfor %}
</VirtualHost>
```

Find more examples for including additional templates, macros, functions, .. [here](https://github.com/flosch/pongo2#pongo2)
and [here](https://docs.djangoproject.com/en/dev/topics/templates/)

Use as library
--------------

tmpl follows ('ish) the [Standard Package Layout, as defined by Ben Johnson](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1),
if that helps. Check out the [example folder](https://github.com/ukautz/tmpl/tree/master/example), the `BuildTmpl`
facade function in [tmpl.go](https://github.com/ukautz/tmpl/blob/master/tmpl.go) and the integration tests in
[imports_test.go](https://github.com/ukautz/tmpl/blob/master/imports/imports_test.go) to get an understanding on how to use.

To use guessers and builders, you can import the whole bundle:

```go
package mypackage

import (
	"fmt"
	"..."
	_ "github.com/ukautz/tmpl/imports"
)

// --- %< ---

renderer, err := tmpl.GuessRenderer("http://some/url.tmpl")  // or "file:///etc/file.pongo" or ..
source, err := tmpl.GuessSource("http://some/url") // or "file:///path" or ..
decoder, err := tmpl.GuessDecoder("http://some/url.json") // or "/srv/path/file.yaml" or ..
```
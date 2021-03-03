# tmpl 

tmpl is a lightweight, shell script friendly document renderer. Something like `envsubst` just different.

tmpl is inspired by recent extensions of the shell toolset like [jq](https://github.com/stedolan/jq), which allows to
work with complex, structured (JSON) data from the shell. This tool provides a solution for a special case in that
particular niche: automated generation of complex documents out of structured data using smart templates.

## Examples

```shell
# hello world
$ tmpl -d 'env:' -t 'envsubst:hello ${USER}, how are you?'

# print out all env vars as JSON
$ tmpl -d 'env:' -t 'template:{{ json .data }}'

# use source from remote location to render local template and store the result
# in a local file
$ tmpl --data-location https://config.acme.tld/node/$(uname -n).json \
    --template-location /etc/apache/templates/homepage.tmpl \
    --output /var/www/home/index.html

# get data from output of execution, then pull template from remote # location
# and pipe output to script that sends alarm
$ tmpl --data-location "shell://elasticdump ... args ..." \
    --decoder json \
    --template-location "https://wiki.acme.tld/templates/report.tmpl?auth=$(tokengen.sh)" | \
    send-alarm.sh ...
```

There are many possible applications or integrations. Here are a few:

* **Reporting**, eg generate HTML reports or summaries based on "native JSON sources" like ElasticSearch or MongoDB
* **Build pipeline**, eg
  * transform generic YAML/JSON into domain specific configuration files (think YAML -> NGINX config)
  * render templates from environment variables
* **Monitoring / alerting**, eg generate system alert emails with complex information rendered in a human readable way
* **Transactional messaging**, eg create rich documents from raw JSON (eg JSON -> DocX)
* you can think of something ..



## Data Locations

A data location contains input data (structures). Some locations indicate a data format, and thereby a Decoder. support guessing of the decoder, e.g. an URL like `http://acme.tld/foo.json` implies JSON format.

**Supported locations** are:

* `env:` or `env:SOME_PREFIX_` converting all (prefix) matching env vars into a flat data map; for example `env:FOO_` converts `FOO_PARAM=x` and `FOO_Other=yy` into `{"PARAM":"x", "Other": "yy"}`; uses JSON internally
* `http://` or `https://`: arbitrary, GETable HTTP(S) URLs; decoder guessed from file ending like `.json` or `.yaml` ending of file in URL path (`https://acme.tld/my/file.json`)
* `file:///path/to/file` or `/path/to/file`: arbitrary local files; decoder guessed from file ending like `.json` or `.yaml`
* `shell://`: arbitrary, atomic shell command lines, eg `shell://date +%F`, which would execute `echo '{"foo":"bar"}'` or anything that would return JSON/YAML on STDOUT. Don't use pipes or somesuch..
* `-`: STDIN, requires decoder specification


### Data Decoder

Supported decoders are JSON and YAML. Per default, tmpl tries to guess the format from the URL. The decoder can be set explicitly with `--decoder <json|yaml>` (or `-d <json|yaml>`).
## Template Locations

tmpl supports multiple template render engines:
- [**template**, from the go standard libraries](https://golang.org/pkg/text/template/)
- [**pongo2** (go implementation)](https://github.com/flosch/pongo2) which implements
[Python Django's templating engine](https://docs.djangoproject.com/en/dev/topics/templates/)
- [**envsubst** (go implementation)](github.com/drone/envsubst), which supports a syntax close to the [envsubst](https://linux.die.net/man/1/envsubst) command line tool

The same as with data sources: specify the renderer explicitly with `--renderer | -r <name>` or let tmpl try to guess
it from the template location:

* Template URLs with `.envsubst` file name endings default to `envsubst` engine. Examples:
  * `http://acme.tld/templates/vhost.envsubst?t=123456`
  * `file:///etc/apache/vhost.envsubst`
* Template URLs with `.tmpl` or `.template` file names default to `template` engine. Examples:
  * `http://acme.tld/templates/vhost.tmpl?foo=bar`
  * `file:///etc/apache/vhost.template`
* Template URLs with `.pongo2` or `.pongo` files name endings default to `pongo2` engine. Examples:
  * `http://acme.tld/templates/vhost.pongo2?t=123456`
  * `file:///etc/apache/vhost.pongo`

For template and pongo2 see the examples below, assume the following data structure & content:

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

### `envsubst` engine

This engine only supports flat data structures and is intended to use with the `env:` data location.

Assuming the following env vars are set:

```shell
USER=myself
HOME=/home/myself
APP_NAME=the-app
APP_DOMAIN=the-domain.tld
```

#### Using all env variables

With the following template

```
Hello ${USER}, here is your home: ${HOME}. Your application is named ${APP_NAME}.
```

executed with `tmpl -d 'env:' -t 'file:///path/to/template'` would render:

```text
Hello myself, here is your home: /home/myself. Your application is named the-app.
```

#### Using prefixed env variables

With the following template

```
App name ${NAME} has domain ${DOMAIN}
```

executed with `tmpl -d 'env:APP_' -t 'file:///path/to/template'` would render:

```text
App name the-app has domain the-domain.tld
```


## Use as library

tmpl is (mostly) structured in the [Standard Go Project Layout](https://github.com/golang-standards/project-layout) and follows ('ish) the [Standard Package Layout, as defined by Ben Johnson](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1).

- **Note**: Parts of the library in [`pkg/`](pkg/) is using singletons, intended to be used in a program.
- **Note**: Check out the [example folder](https://github.com/ukautz/tmpl/tree/master/example), the `BuildTmpl` facade function in [tmpl.go](pkg/tmpl.go) and the integration tests in [imports_test.go](pkg/imports/imports_test.go) to get an understanding on how to use.

To use guessers and builders, you can import the whole bundle:

```go
package mypackage

import (
	"fmt"
	"..."
	_ "github.com/ukautz/tmpl/pkg/imports"
)

// --- %< ---

renderer, err := tmpl.GuessRenderer("http://some/url.tmpl")  // or "file:///etc/file.pongo" or ..
source, err := tmpl.GuessSource("http://some/url") // or "file:///path" or ..
decoder, err := tmpl.GuessDecoder("http://some/url.json") // or "/srv/path/file.yaml" or ..
```
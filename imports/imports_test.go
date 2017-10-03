package imports

import (
	"github.com/stretchr/testify/assert"
	"github.com/ukautz/tmpl"
	_ "github.com/ukautz/tmpl"
	_ "github.com/ukautz/tmpl/http"
	_ "github.com/ukautz/tmpl/json"
	_ "github.com/ukautz/tmpl/pongo2"
	_ "github.com/ukautz/tmpl/shellwords"
	_ "github.com/ukautz/tmpl/template"
	_ "github.com/ukautz/tmpl/yaml"
	"reflect"
	"testing"
)

func TestDecoderImports(t *testing.T) {
	for _, name := range []string{"json", "yaml"} {
		decoder, err := tmpl.BuildDecoder(name)
		assert.Nil(t, err, "Build decoder \"%s\" should work", name)
		assert.NotNil(t, decoder, "Build decoder \"%s\" should be created", name)
	}
	for location, typ := range map[string]string{
		"/etc/foo.json": "*json.Decoder",
		"/etc/foo.js":   "*json.Decoder",
		"/etc/foo.yaml": "*yaml.Decoder",
		"/etc/foo.yml":  "*yaml.Decoder",
	} {
		decoder, err := tmpl.GuessDecoder(location)
		assert.Nil(t, err, "Guess decoder \"%s\" should work", location)
		assert.NotNil(t, decoder, "Guess decoder \"%s\" should be created", location)
		rtyp := reflect.TypeOf(decoder).String()
		assert.Equal(t, typ, rtyp, "Guessed decoder \"%s\" should be \"%s\" but is \"%s\"", location, typ, rtyp)
	}
}

func TestRendererImports(t *testing.T) {
	for _, name := range []string{"template", "pongo2"} {
		renderer, err := tmpl.BuildRenderer(name)
		assert.Nil(t, err, "Build renderer \"%s\" should work", name)
		assert.NotNil(t, renderer, "Build renderer \"%s\" should be created", name)
	}
	for location, typ := range map[string]string{
		"/etc/foo.template": "*template.Renderer",
		"/etc/foo.tmpl":     "*template.Renderer",
		"/etc/foo.pongo":    "*pongo2.Renderer",
		"/etc/foo.pongo2":   "*pongo2.Renderer",
	} {
		renderer, err := tmpl.GuessRenderer(location)
		assert.Nil(t, err, "Guess renderer \"%s\" should work", location)
		assert.NotNil(t, renderer, "Guess renderer \"%s\" should be created", location)
		rtyp := reflect.TypeOf(renderer).String()
		assert.Equal(t, typ, rtyp, "Guessed renderer \"%s\" should be \"%s\" but is \"%s\"", location, typ, rtyp)
	}
}

func TestSourceImports(t *testing.T) {
	for location, typ := range map[string]string{
		"http://acme.tld":    "http.Source",
		"shell://ps auxf":    "shellwords.Source",
		"file:///etc/config": "tmpl.FileSource",
		"/etc/config":        "tmpl.FileSource",
		"-":                  "*tmpl.InlineSource",
	} {
		source, err := tmpl.GuessSource(location)
		assert.Nil(t, err, "Guess source \"%s\" should work", location)
		assert.NotNil(t, source, "Guess source \"%s\" should be created", location)
		rtyp := reflect.TypeOf(source).String()
		assert.Equal(t, typ, rtyp, "Guessed source \"%s\" should be \"%s\" but is \"%s\"", location, typ, rtyp)
	}
}

package imports

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	tmpl "github.com/ukautz/tmpl/pkg"
)

func TestDecoderImports(t *testing.T) {
	for _, name := range []string{"json", "yaml"} {
		decoder, err := tmpl.BuildDecoder(name)
		assert.Nil(t, err, "Build decoder \"%s\" should work", name)
		assert.NotNil(t, decoder, "Build decoder \"%s\" should be created", name)
	}

	for location, typ := range map[string]string{
		"/etc/foo.json": "tmpl.JSONDecoder",
		"/etc/foo.js":   "tmpl.JSONDecoder",
		"/etc/foo.yaml": "yaml.Decoder",
		"/etc/foo.yml":  "yaml.Decoder",
		"env:anything":  "tmpl.JSONDecoder",
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
		"/etc/foo.template": "tmpl.TemplateRenderer",
		"/etc/foo.tmpl":     "tmpl.TemplateRenderer",
		"/etc/foo.pongo":    "pongo2.Renderer",
		"/etc/foo.pongo2":   "pongo2.Renderer",
		"/etc/foo.envsubst": "envsubst.Renderer",
	} {
		renderer, err := tmpl.GuessRenderer(location)
		assert.Nil(t, err, "Guess renderer \"%s\" should work", location)
		assert.NotNil(t, renderer, "Guess renderer \"%s\" should be created", location)
		rtyp := reflect.TypeOf(renderer).String()
		assert.Equal(t, typ, rtyp, "Guessed renderer \"%s\" should be \"%s\" but is \"%s\"", location, typ, rtyp)
	}
}

func TestLocationImports(t *testing.T) {
	for location, typ := range map[string]string{
		"http://acme.tld":                        "tmpl.HTTPLocation",
		"shell://ps auxf":                        "shellwords.Location",
		"file:///etc/config":                     "tmpl.FileLocation",
		filepath.Join("fixtures", "config-file"): "tmpl.FileLocation",
		"template:foo":                           "tmpl.TemplateRawLocation",
		"pongo:foo":                              "tmpl.TemplateRawLocation",
		"pongo2:foo":                             "tmpl.TemplateRawLocation",
		"envsubst:foo":                           "tmpl.TemplateRawLocation",
		"env:FOO_":                               "tmpl.OSEnvLocation",
		"-":                                      "tmpl.STDINLocation",
	} {
		loc, err := tmpl.GuessLocation(location)
		assert.Nil(t, err, "Guess location for \"%s\" should work", location)
		assert.NotNil(t, loc, "Guess location \"%s\" should be created", location)
		rtyp := reflect.TypeOf(loc).String()
		assert.Equal(t, typ, rtyp, "Guessed location \"%s\" should be \"%s\" but is \"%s\"", location, typ, rtyp)
	}
}

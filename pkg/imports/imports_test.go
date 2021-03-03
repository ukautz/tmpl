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

func TestSourceImports(t *testing.T) {
	for location, typ := range map[string]string{
		"http://acme.tld":                        "tmpl.HTTPSource",
		"shell://ps auxf":                        "shellwords.Source",
		"file:///etc/config":                     "tmpl.FileSource",
		filepath.Join("fixtures", "config-file"): "tmpl.FileSource",
		"template:foo":                           "tmpl.TemplateRawSource",
		"pongo:foo":                              "tmpl.TemplateRawSource",
		"pongo2:foo":                             "tmpl.TemplateRawSource",
		"envsubst:foo":                           "tmpl.TemplateRawSource",
		"env:FOO_":                               "tmpl.OSEnvSource",
		"-":                                      "tmpl.STDINSource",
	} {
		source, err := tmpl.GuessSource(location)
		assert.Nil(t, err, "Guess source for \"%s\" should work", location)
		assert.NotNil(t, source, "Guess source \"%s\" should be created", location)
		rtyp := reflect.TypeOf(source).String()
		assert.Equal(t, typ, rtyp, "Guessed source \"%s\" should be \"%s\" but is \"%s\"", location, typ, rtyp)
	}
}

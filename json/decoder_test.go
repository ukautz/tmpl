package json

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"reflect"
	"fmt"
)

func TestNewDecoder(t *testing.T) {
	d := NewDecoder()
	assert.NotNil(t, d, "Decoder created")
	assert.Equal(t, "*json.Decoder", reflect.TypeOf(d).String(), "Decoder expected")
}

func TestBuildDecoder(t *testing.T) {
	ok := []string{
		"http://acme.tld/data.json",
		"http://acme.tld/data.js",
		"/path/to/data.js",
	}
	notOk := []string{
		"http://acme.tld/data.yaml",
		"http://acme.tld/data.yml",
		"/path/to/data.yml",
	}
	for _, loc := range ok {
		d := BuildDecoder(loc)
		assert.NotNil(t, d, fmt.Sprintf("Decoder created from %s", loc))
	}
	for _, loc := range notOk {
		d := BuildDecoder(loc)
		assert.Nil(t, d, fmt.Sprintf("Decoder NOT created from %s", loc))
	}
}

func TestDecoder_Decode(t *testing.T) {
	d := &Decoder{}

	res, err := d.Decode([]byte(`{"foo": "bar", "baz": ["zoing", "zing"]}`))
	assert.Nil(t, err, "Should not error")
	assert.Equal(t, res, map[string]interface{}{"foo": "bar", "baz": []interface{}{"zoing", "zing"}}, "Should parse into map")

	res, err = d.Decode([]byte(`["foo", "bar", {"baz": "zoing"}]`))
	assert.Nil(t, err, "Should not error")
	assert.Equal(t, res, []interface{}{"foo", "bar", map[string]interface{}{"baz": "zoing"}}, "Should parse into slice")

	res, err = d.Decode([]byte(`"foo"`))
	assert.NotNil(t, err, "Should throw error")
	assert.Equal(t, err.Error(), "could not decode JSON data", "Should not recognize")
}
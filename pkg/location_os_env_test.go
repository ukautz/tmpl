package tmpl

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildOSEnvLocation(t *testing.T) {
	assert.NotNil(t, BuildOSEnvLocation("env:"))
	assert.NotNil(t, BuildOSEnvLocation("env:FOO"))
	assert.NotNil(t, BuildOSEnvLocation("env:__bla23123"))
	assert.NotNil(t, BuildOSEnvLocation("env:X_Gjas8752"))
	assert.Nil(t, BuildOSEnvLocation("env"))
	assert.Nil(t, BuildOSEnvLocation(""))
	assert.Nil(t, BuildOSEnvLocation("file://"))
}

func TestOSEnvLocation_Load(t *testing.T) {
	prefix := "__GO_TEST__"
	for k, v := range map[string]string{
		"_foo":           "foo1",
		"__bar":          "bar1",
		"____baz":        "baz1",
		prefix + "_foo":  "FOO2",
		prefix + "__bar": "BAR2",
	} {
		os.Setenv(k, v)
	}

	env := OSEnvLocation(prefix)
	raw, err := env.Load()
	assert.Nil(t, err)
	assert.JSONEq(t, `{"_foo":"FOO2","__bar":"BAR2"}`, string(raw))
}

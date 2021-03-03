package tmpl

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildOSEnvSource(t *testing.T) {
	assert.NotNil(t, BuildOSEnvSource("env:"))
	assert.NotNil(t, BuildOSEnvSource("env:FOO"))
	assert.NotNil(t, BuildOSEnvSource("env:__bla23123"))
	assert.NotNil(t, BuildOSEnvSource("env:X_Gjas8752"))
	assert.Nil(t, BuildOSEnvSource("env"))
	assert.Nil(t, BuildOSEnvSource(""))
	assert.Nil(t, BuildOSEnvSource("file://"))
}

func TestOSEnvSource_Load(t *testing.T) {
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

	env := OSEnvSource(prefix)
	raw, err := env.Load()
	assert.Nil(t, err)
	assert.JSONEq(t, `{"_foo":"FOO2","__bar":"BAR2"}`, string(raw))
}

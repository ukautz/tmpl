package tmpl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildTemplateRawSource(t *testing.T) {
	assert.NotNil(t, BuildTemplateRawSource("template:"))
	assert.NotNil(t, BuildTemplateRawSource("template:a-template"))
	assert.NotNil(t, BuildTemplateRawSource("pongo:"))
	assert.NotNil(t, BuildTemplateRawSource("pongo:a-template"))
	assert.NotNil(t, BuildTemplateRawSource("pongo2:"))
	assert.NotNil(t, BuildTemplateRawSource("pongo2:a-template"))
	assert.NotNil(t, BuildTemplateRawSource("envsubst:"))
	assert.NotNil(t, BuildTemplateRawSource("envsubst:a-template"))
	assert.Nil(t, BuildTemplateRawSource(""))
	assert.Nil(t, BuildTemplateRawSource("input"))
}

func TestTemplateRawSource_Load(t *testing.T) {
	input := TemplateRawSource(`the-input`)
	raw, err := input.Load()
	assert.Nil(t, err)
	assert.Equal(t, `the-input`, string(raw))
}

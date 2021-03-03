package tmpl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildTemplateRawLocation(t *testing.T) {
	assert.NotNil(t, BuildTemplateRawLocation("template:"))
	assert.NotNil(t, BuildTemplateRawLocation("template:a-template"))
	assert.NotNil(t, BuildTemplateRawLocation("pongo:"))
	assert.NotNil(t, BuildTemplateRawLocation("pongo:a-template"))
	assert.NotNil(t, BuildTemplateRawLocation("pongo2:"))
	assert.NotNil(t, BuildTemplateRawLocation("pongo2:a-template"))
	assert.NotNil(t, BuildTemplateRawLocation("envsubst:"))
	assert.NotNil(t, BuildTemplateRawLocation("envsubst:a-template"))
	assert.Nil(t, BuildTemplateRawLocation(""))
	assert.Nil(t, BuildTemplateRawLocation("input"))
}

func TestTemplateRawLocation_Load(t *testing.T) {
	input := TemplateRawLocation(`the-input`)
	raw, err := input.Load()
	assert.Nil(t, err)
	assert.Equal(t, `the-input`, string(raw))
}

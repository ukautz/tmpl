package shellwords

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"runtime"
)

func TestBuildSource(t *testing.T) {
	s := BuildSource("shell://arbitrary string")
	assert.NotNil(t, s, "Source should build from shell schema")
	assert.Equal(t, []string{"arbitrary", "string"}, []string(s.(Source)), "Should result into split string")
	s = BuildSource("http://valid.tld/")
	assert.Nil(t, s, "Source should NOT build from HTTP URL")
	s = BuildSource("file:///valid/path")
	assert.Nil(t, s, "Source should NOT build from file URL")
}

func TestSource_Load(t *testing.T) {
	s, err := NewSource("go version")
	assert.Nil(t, err, "No creating error")
	assert.NotNil(t, s, "Created")
	res, err := s.Load()
	assert.Nil(t, err, "No load error")
	assert.Contains(t, string(res), runtime.Version(), "Expected output")
}
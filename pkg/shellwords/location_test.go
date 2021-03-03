package shellwords

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildLocation(t *testing.T) {
	s := BuildLocation("shell://arbitrary string")
	assert.NotNil(t, s, "Location should build from shell schema")
	assert.Equal(t, []string{"arbitrary", "string"}, []string(s.(Location)), "Should result into split string")
	s = BuildLocation("http://valid.tld/")
	assert.Nil(t, s, "Location should NOT build from HTTP URL")
	s = BuildLocation("file:///valid/path")
	assert.Nil(t, s, "Location should NOT build from file URL")
}

func TestLocation_Load(t *testing.T) {
	s, err := NewLocation("go version")
	assert.Nil(t, err, "No creating error")
	assert.NotNil(t, s, "Created")
	res, err := s.Load()
	assert.Nil(t, err, "No load error")
	assert.Contains(t, string(res), runtime.Version(), "Expected output")
}

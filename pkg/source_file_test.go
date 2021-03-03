package tmpl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildFileSource(t *testing.T) {
	ok := []string{
		"file:///foo.txt",
		// TODO: "/foo.txt",
	}
	notOk := []string{
		"-",
		"http://foo.tld/",
		"shell://foo.tld/",
		"/foo\n.txt",
	}
	for _, loc := range ok {
		assert.NotNil(t, BuildFileSource(loc), fmt.Sprintf("File source build from \"%s\"", loc))
	}
	for _, loc := range notOk {
		assert.Nil(t, BuildFileSource(loc), fmt.Sprintf("File source NOT build from \"%s\"", loc))
	}
}

func TestFileSource_Load(t *testing.T) {
	fs := FileSource("./source.go")
	raw, err := fs.Load()
	assert.Nil(t, err, "No load error")
	assert.Contains(t, string(raw), "package tmpl", "File was loaded")
}

func TestFileSource_Load_FromFileSchema(t *testing.T) {
	fs := FileSource("file://./source.go")
	raw, err := fs.Load()
	assert.Nil(t, err, "No load error")
	assert.Contains(t, string(raw), "package tmpl", "File was loaded")
}

func TestFileSource_Load_FailFromNoFile(t *testing.T) {
	fs := FileSource("not-existing-file")
	_, err := fs.Load()
	assert.NotNil(t, err, "No throw error")
	assert.Contains(t, err.Error(), "open not-existing-file: no such file or directory", "Should remark does not exists")
}

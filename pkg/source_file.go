package tmpl

import (
	"io/ioutil"
	"regexp"
	"strings"
)

// FileSource loads local files; identified by prefix "file://"
type FileSource string

func (s FileSource) Load() ([]byte, error) {
	path := string(s)
	if strings.HasPrefix(strings.ToLower(path), "file://") {
		path = path[len("file://"):]
	}
	return ioutil.ReadFile(path)
}

var fileSourceHasSchema = regexp.MustCompile(`^[a-zA-Z0-9]+://`)

// BuildFileSource returns FileSource if location is local file (with or without file:// schema)
func BuildFileSource(location string) Source {
	if strings.HasPrefix(strings.ToLower(location), "file://") {
		return FileSource(location)
	}
	return nil
}

func init() {
	Sources = append(Sources, BuildFileSource)
}

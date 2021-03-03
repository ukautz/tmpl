package tmpl

import (
	"io/ioutil"
	"regexp"
	"strings"
)

// FileLocation loads local files; identified by prefix "file://"
type FileLocation string

func (s FileLocation) Load() ([]byte, error) {
	path := string(s)
	if strings.HasPrefix(strings.ToLower(path), "file://") {
		path = path[len("file://"):]
	}
	return ioutil.ReadFile(path)
}

var fileLocationHasSchema = regexp.MustCompile(`^[a-zA-Z0-9]+://`)

// BuildFileLocation returns FileLocation if location is local file (with or without file:// schema)
func BuildFileLocation(location string) Location {
	if strings.HasPrefix(strings.ToLower(location), "file://") {
		return FileLocation(location)
	}
	return nil
}

func init() {
	Locations = append(Locations, BuildFileLocation)
}

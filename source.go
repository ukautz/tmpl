package tmpl

import (
	"errors"
	"io/ioutil"
	"strings"
	"regexp"
	"os"
)

type (

	// Source is interface for any loadable data/resource/bytes location
	Source interface {
		// Load
		Load() ([]byte, error)
	}

	// FileSource loads local files
	FileSource string

	// InlineSource loads from STDIN
	InlineSource struct{}
)

var Sources = []func(string) Source{}
var hasSchema = regexp.MustCompile(`^[a-zA-Z0-9]+://`)


func (is *InlineSource) Load() ([]byte, error) {
	return ioutil.ReadAll(os.Stdin)
}

func (fs FileSource) Load() ([]byte, error) {
	path := string(fs)
	if strings.HasPrefix(strings.ToLower(path), "file://") {
		path = path[len("file://"):]
	}
	return ioutil.ReadFile(path)
}

// GuessSource tries to "best applicable source" for provided location (URL, path, ..) or returns error
func GuessSource(location string) (Source, error) {
	for _, builder := range Sources {
		if source := builder(location); source != nil {
			return source, nil
		}
	}

	return nil, errors.New("no source accepting given location registered")
}

// BuildInlineSource returns InlineSource if location is "-"
func BuildInlineSource(location string) Source {
	if location == "-" {
		return &InlineSource{}
	}
	return nil
}

// BuildFileSource returns FileSource if location is local file (with or without file:// schema)
func BuildFileSource(location string) Source {
	if strings.Contains(location, "\n") || strings.Contains(location, "\r") {
		return nil
	} else if strings.HasPrefix(strings.ToLower(location), "file://") || !(hasSchema.MatchString(location) || location == "-") {
		return FileSource(location)
	}
	return nil
}

func init() {
	Sources = append(Sources, BuildFileSource, BuildInlineSource)
}
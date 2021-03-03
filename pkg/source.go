package tmpl

import (
	"fmt"
	"os"
	"strings"
)

// Source is interface for any loadable data/resource/bytes location
type Source interface {
	// Load
	Load() ([]byte, error)
}

var Sources = []func(string) Source{}

// GuessSource tries to "best applicable source" for provided location (URL, path, ..) or returns error
func GuessSource(location string) (Source, error) {
	for _, builder := range Sources {
		if source := builder(location); source != nil {
			return source, nil
		}
	}

	if !strings.Contains(location, "\n") && !strings.Contains(location, "\n") {
		_, err := os.Stat(location)
		if err == nil {
			return FileSource(location), nil
		}
	}

	return nil, fmt.Errorf("no source for `%s` registered", location)
}

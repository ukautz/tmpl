package tmpl

import (
	"fmt"
	"os"
	"strings"
)

// Location is interface for any loadable data/resource/bytes location
type Location interface {
	// Load
	Load() ([]byte, error)
}

var Locations = []func(string) Location{}

// GuessLocation tries to find "best applicable" Location for provided <URL|path|..>
func GuessLocation(location string) (Location, error) {
	for _, builder := range Locations {
		if loc := builder(location); loc != nil {
			return loc, nil
		}
	}

	if !strings.Contains(location, "\n") && !strings.Contains(location, "\r") {
		_, err := os.Stat(location)
		if err == nil {
			return FileLocation(location), nil
		}
	}

	return nil, fmt.Errorf("no location for `%s` registered", location)
}

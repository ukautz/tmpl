package tmpl

import (
	"encoding/json"
	"os"
	"strings"
)

// OSEnvLocation loads from environment variables; identified by prefix "env:<arbitrary_prefix_or_empty>"; only for
type OSEnvLocation string

func (s OSEnvLocation) Load() ([]byte, error) {
	prefix := string(s)
	strip := len(prefix)
	result := make(map[string]string)
	for _, name := range os.Environ() {
		name = name[:strings.Index(name, "=")]
		if strings.HasPrefix(name, prefix) {
			result[name[strip:]] = os.Getenv(name)
		}
	}
	return json.Marshal(result)
}

// BuildOSEnvLocation returns OSEnvLocation if location is local file (with or without file:// schema)
func BuildOSEnvLocation(location string) Location {
	if strings.HasPrefix(location, "env:") {
		return OSEnvLocation(location[4:])
	}
	return nil
}

func init() {
	Locations = append(Locations, BuildOSEnvLocation)
}

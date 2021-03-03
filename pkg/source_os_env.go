package tmpl

import (
	"encoding/json"
	"os"
	"strings"
)

// OSEnvSource loads from environment variables; identified by prefix "env:<arbitrary_prefix_or_empty>"; only for
type OSEnvSource string

func (s OSEnvSource) Load() ([]byte, error) {
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

// BuildFileSource returns FileSource if location is local file (with or without file:// schema)
func BuildOSEnvSource(location string) Source {
	if strings.HasPrefix(location, "env:") {
		return OSEnvSource(location[4:])
	}
	return nil
}

func init() {
	Sources = append(Sources, BuildOSEnvSource)
}

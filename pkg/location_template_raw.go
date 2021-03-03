package tmpl

import (
	"strings"
)

// TemplateRawLocation loads from parameter; identified by "input:"; intended _only_ for template
type TemplateRawLocation string

func (s TemplateRawLocation) Load() ([]byte, error) {
	return []byte(s), nil
}

var templateRawPrefices = []string{
	"envsubst",
	"pongo",
	"pongo2",
	"template",
}

func BuildTemplateRawLocation(location string) Location {
	for _, prefix := range templateRawPrefices {
		if strings.HasPrefix(location, prefix+":") {
			return TemplateRawLocation(location[len(prefix)+1:])
		}
	}
	return nil
}

func init() {
	Locations = append(Locations, BuildTemplateRawLocation)
}

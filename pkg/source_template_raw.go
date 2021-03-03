package tmpl

import (
	"strings"
)

// TemplateRawSource loads from parameter; identified by "input:"; intended _only_ for template
type TemplateRawSource string

func (s TemplateRawSource) Load() ([]byte, error) {
	return []byte(s), nil
}

var templateRawPrefices = []string{
	"envsubst",
	"pongo",
	"pongo2",
	"template",
}

func BuildTemplateRawSource(location string) Source {
	for _, prefix := range templateRawPrefices {
		if strings.HasPrefix(location, prefix+":") {
			return TemplateRawSource(location[len(prefix)+1:])
		}
	}
	return nil
}

func init() {
	Sources = append(Sources, BuildTemplateRawSource)
}

package tmpl

import (
	"io/ioutil"
	"os"
)

// STDINSource loads from STDIN; identified by "-"; intended _only_ for
type STDINSource struct{}

func (s STDINSource) Load() ([]byte, error) {
	return ioutil.ReadAll(os.Stdin)
}

func NewSTDINSource() STDINSource {
	return STDINSource{}
}

// BuildSTDINSource returns InlineSource if location is "-"
func BuildSTDINSource(location string) Source {
	if location == "-" {
		return STDINSource{}
	}
	return nil
}

func init() {
	Sources = append(Sources, BuildSTDINSource)
}

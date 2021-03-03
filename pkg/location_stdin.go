package tmpl

import (
	"io/ioutil"
	"os"
)

// STDINLocation loads from STDIN; identified by "-"; intended _only_ for
type STDINLocation struct{}

func (s STDINLocation) Load() ([]byte, error) {
	return ioutil.ReadAll(os.Stdin)
}

func NewSTDINLocation() STDINLocation {
	return STDINLocation{}
}

// BuildSTDINLocation returns STDINLocation if location is "-"
func BuildSTDINLocation(location string) Location {
	if location == "-" {
		return STDINLocation{}
	}
	return nil
}

func init() {
	Locations = append(Locations, BuildSTDINLocation)
}

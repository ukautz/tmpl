package yaml

import (
	"errors"
	"regexp"

	tmpl "github.com/ukautz/tmpl/pkg"
	"gopkg.in/yaml.v2"
)

type Decoder struct{}

var isYamlLocation = regexp.MustCompile(`\.ya?ml(?:$|\?)`)

func (d Decoder) Decode(data []byte) (interface{}, error) {
	var a []interface{}
	var m map[string]interface{}
	if err := yaml.Unmarshal(data, &a); err == nil {
		return a, nil
	} else if err = yaml.Unmarshal(data, &m); err == nil {
		return m, nil
	}
	return nil, errors.New("could not decode JSON data")
}

func NewDecoder() tmpl.Decoder {
	return Decoder{}
}

func BuildDecoder(location string) tmpl.Decoder {
	if isYamlLocation.MatchString(location) {
		return NewDecoder()
	}
	return nil
}

func init() {
	tmpl.Decoders["yaml"] = NewDecoder
	tmpl.DecoderGuesses = append(tmpl.DecoderGuesses, BuildDecoder)
}

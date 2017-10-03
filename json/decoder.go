package json

import (
	"encoding/json"
	"errors"
	"github.com/ukautz/tmpl"
	"regexp"
)

type Decoder struct{}

var isJsonLocation = regexp.MustCompile(`\.js(?:on)?(?:$|\?)`)

func (d *Decoder) Decode(data []byte) (interface{}, error) {
	var a []interface{}
	var m map[string]interface{}
	if err := json.Unmarshal(data, &a); err == nil {
		return a, nil
	} else if err = json.Unmarshal(data, &m); err == nil {
		return m, nil
	}
	return nil, errors.New("could not decode JSON data")
}

func NewDecoder() tmpl.Decoder {
	return &Decoder{}
}

func BuildDecoder(location string) tmpl.Decoder {
	if isJsonLocation.MatchString(location) {
		return &Decoder{}
	}
	return nil
}

func init() {
	tmpl.Decoders["json"] = NewDecoder
	tmpl.DecoderGuesses = append(tmpl.DecoderGuesses, BuildDecoder)
}
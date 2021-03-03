package tmpl

import (
	"encoding/json"
	"errors"
	"regexp"
)

type JSONDecoder struct{}

func (d JSONDecoder) Decode(data []byte) (interface{}, error) {
	var a []interface{}
	var m map[string]interface{}
	if err := json.Unmarshal(data, &a); err == nil {
		return a, nil
	} else if err = json.Unmarshal(data, &m); err == nil {
		return m, nil
	}
	return nil, errors.New("could not decode JSON data")
}

func NewJSONDecoder() Decoder {
	return JSONDecoder{}
}

var isJsonLocation = regexp.MustCompile(`(?:^env:|\.js(?:on)?(?:$|\?))`)

func BuildJSONDecoder(location string) Decoder {
	if isJsonLocation.MatchString(location) {
		return NewJSONDecoder()
	}
	return nil
}

func init() {
	Decoders["json"] = NewJSONDecoder
	DecoderGuesses = append(DecoderGuesses, BuildJSONDecoder)
}

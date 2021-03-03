package tmpl

import (
	"errors"
	"fmt"
)

type (
	Decoder interface {
		Decode(data []byte) (interface{}, error)
	}
)

var Decoders = map[string]func() Decoder{}
var DecoderGuesses = []func(location string) Decoder{}

func GuessDecoder(location string) (Decoder, error) {
	for _, builder := range DecoderGuesses {
		if decoder := builder(location); decoder != nil {
			return decoder, nil
		}
	}
	return nil, errors.New("could not guess decoder for location")
}

func BuildDecoder(name string) (Decoder, error) {
	if builder, ok := Decoders[name]; ok {
		return builder(), nil
	}
	return nil, fmt.Errorf("no decoder with name \"%s\" registered", name)
}

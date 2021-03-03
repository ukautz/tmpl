package tmpl

import (
	"fmt"
)

// Renderer is interface for template rendering engine adapters. Engines process data (map[string]interface{} or
// []interfaces{}) using provided template into combined result.
//
type Renderer interface {
	Render(data interface{}, template []byte) ([]byte, error)
}

var Renderers = map[string]func() Renderer{}
var RendererGuesses = []func(location string) Renderer{}

// GuessRenderer tries to determine
func GuessRenderer(location string) (Renderer, error) {
	for _, builder := range RendererGuesses {
		if renderer := builder(location); renderer != nil {
			return renderer, nil
		}
	}
	return nil, fmt.Errorf("could not guess renderer for location %s", location)
}

func BuildRenderer(name string) (Renderer, error) {
	if builder, ok := Renderers[name]; ok {
		return builder(), nil
	}
	return nil, fmt.Errorf("no rendered with name \"%s\" registered", name)
}

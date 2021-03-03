// Package tmpl implements template generation and variable substition from a large set of data sources
package tmpl

import (
	"errors"
	"fmt"
)

// Tmpl is a facade for the template generation process. It creates (partially best guesses)
type Tmpl struct {
	Data     Source
	Decoder  Decoder
	Renderer Renderer
	Template Source
}

// Build constructs the Tmpl facade from names of components
func Build(dataSource, decoderName, templateSource, rendererName string) (*Tmpl, error) {
	if dataSource == "-" && dataSource == templateSource {
		return nil, errors.New("cannot use STDIN for data and template at the same time")
	}

	data, err := GuessSource(dataSource)
	if err != nil {
		return nil, fmt.Errorf("failed to create data source: %w", err)
	}

	var decoder Decoder
	if decoderName == "guess" {
		decoder, err = GuessDecoder(dataSource)
	} else {
		decoder, err = BuildDecoder(decoderName)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to create data decoder: %w", err)
	}

	template, err := GuessSource(templateSource)
	if err != nil {
		return nil, fmt.Errorf("failed to create data location: %w", err)
	}

	var renderer Renderer
	if rendererName == "guess" {
		renderer, err = GuessRenderer(templateSource)
	} else {
		renderer, err = BuildRenderer(rendererName)
	}
	if err != nil {
		return nil, err
	}

	return &Tmpl{
		Data:     data,
		Decoder:  decoder,
		Template: template,
		Renderer: renderer,
	}, nil
}

// Produce reads & decodes the data, loads the template and returns the rendered result
func (q *Tmpl) Produce() ([]byte, error) {
	data, err := q.Data.Load()
	if err != nil {
		return nil, err
	}

	decoded, err := q.Decoder.Decode(data)
	if err != nil {
		return nil, err
	}

	template, err := q.Template.Load()
	if err != nil {
		return nil, err
	}

	return q.Renderer.Render(decoded, template)
}

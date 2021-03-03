package pongo2

import (
	"fmt"
	"regexp"

	"github.com/flosch/pongo2/v4"
	tmpl "github.com/ukautz/tmpl/pkg"
)

type Renderer struct {
}

var isPongo2Renderer = regexp.MustCompile(`(?:^pongo2?:|\.pongo2?(?:$|\?))`)

func (r Renderer) Render(data interface{}, template []byte) ([]byte, error) {
	tmpl, err := pongo2.FromString(string(template))
	if err != nil {
		return nil, fmt.Errorf("could not read template: %s", err)
	}

	out, err := tmpl.Execute(pongo2.Context{"data": data})
	if err != nil {
		return nil, fmt.Errorf("could not render template: %s", err)
	}

	return []byte(out), nil
}

func NewRenderer() tmpl.Renderer {
	return Renderer{}
}

func BuildRenderer(location string) tmpl.Renderer {
	if isPongo2Renderer.MatchString(location) {
		return NewRenderer()
	}
	return nil
}

func init() {
	tmpl.Renderers["pongo2"] = NewRenderer
	tmpl.RendererGuesses = append(tmpl.RendererGuesses, BuildRenderer)
}

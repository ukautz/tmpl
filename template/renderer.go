package template

import (
	"fmt"
	"github.com/ukautz/tmpl"
	"text/template"
	"bytes"
	"regexp"
)

type Renderer struct {
}

var isTemplateRenderer = regexp.MustCompile(`\.(?:template|tmpl)(?:$|\?)`)

func (r *Renderer) Render(data interface{}, templateData []byte) ([]byte, error) {
	t, err := template.New("template").Parse(string(templateData))
	if err != nil {
		return nil, fmt.Errorf("could not parse template: %s", err)
	}
	out := bytes.NewBuffer(nil)
	err = t.Execute(out, map[string]interface{}{"data":data})
	if err != nil {
		return nil, fmt.Errorf("could not render template: %s", err)
	}

	return out.Bytes(), nil
}

func NewRenderer() tmpl.Renderer {
	return &Renderer{}
}

func BuildRenderer(location string) tmpl.Renderer {
	if isTemplateRenderer.MatchString(location) {
		return &Renderer{}
	}
	return nil
}

func init() {
	tmpl.Renderers["template"] = NewRenderer
	tmpl.RendererGuesses = append(tmpl.RendererGuesses, BuildRenderer)
}


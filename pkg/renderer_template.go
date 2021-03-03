package tmpl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"text/template"
)

type TemplateRenderer struct {
}

var isTemplateRenderer = regexp.MustCompile(`(?:^template:|\.(?:template|tmpl)(?:$|\?))`)

func (r TemplateRenderer) Render(data interface{}, templateData []byte) ([]byte, error) {
	t, err := template.New("template").Funcs(makeFunctions()).Parse(string(templateData))
	if err != nil {
		return nil, fmt.Errorf("could not parse template: %s", err)
	}
	out := bytes.NewBuffer(nil)
	err = t.Execute(out, map[string]interface{}{"data": data})
	if err != nil {
		return nil, fmt.Errorf("could not render template: %s", err)
	}

	return out.Bytes(), nil
}

func makeFunctions() template.FuncMap {
	return template.FuncMap{
		"json": func(in interface{}) string {
			raw, err := json.Marshal(in)
			if err != nil {
				return err.Error()
			}
			return string(raw)
		},
		"jsonPretty": func(in interface{}) string {
			raw, err := json.MarshalIndent(in, "", "  ")
			if err != nil {
				return err.Error()
			}
			return string(raw)
		},
	}
}

func NewTemplateRenderer() Renderer {
	return TemplateRenderer{}
}

func BuildTemplateRenderer(location string) Renderer {
	if isTemplateRenderer.MatchString(location) {
		return NewTemplateRenderer()
	}
	return nil
}

func init() {
	Renderers["template"] = NewTemplateRenderer
	RendererGuesses = append(RendererGuesses, BuildTemplateRenderer)
}

package envsubst

import (
	"fmt"
	"log"
	"reflect"
	"regexp"

	"github.com/drone/envsubst"
	tmpl "github.com/ukautz/tmpl/pkg"
)

type Renderer struct {
}

var isEnvsubstRenderer = regexp.MustCompile(`(?:^env(?:subst)?:|\.envsubst(?:$|\?))`)

func (r Renderer) Render(data interface{}, templateData []byte) ([]byte, error) {
	strMap, err := toMapStringString(data)
	if err != nil {
		return nil, err
	}

	template, err := envsubst.Eval(string(templateData), func(key string) string {
		log.Printf("TEST REPLACE %s", key)
		if val, ok := strMap[key]; ok {
			return val
		}
		return fmt.Sprintf("<NOTFOUND:%s>", key)
	})

	return []byte(template), nil
}

func toMapStringString(from interface{}) (map[string]string, error) {
	ref := reflect.Indirect(reflect.ValueOf(from))
	if ref.Kind() != reflect.Map {
		return nil, fmt.Errorf("unsupported data kind %s", ref.Kind())
	}

	res := make(map[string]string)
	for _, key := range ref.MapKeys() {
		strKey, err := asString(key)
		if err != nil {
			return nil, fmt.Errorf("invalid key: %w", err)
		}

		strVal, err := asString(ref.MapIndex(key))
		if err != nil {
			return nil, fmt.Errorf("invalid value: %w", err)
		}

		res[strKey] = strVal
	}

	return res, nil
}

func asString(v reflect.Value) (string, error) {
	v = reflect.Indirect(v)
	switch v.Kind() {
	case reflect.String:
		return v.String(), nil
	case reflect.Interface:
		str, ok := v.Interface().(string)
		if ok {
			return str, nil
		}
		return "", fmt.Errorf("unsupported interface %v", v.Interface())
	}
	return "", fmt.Errorf("unsupported kind %s", v.Kind())
}

func NewRenderer() tmpl.Renderer {
	return Renderer{}
}

func BuildRenderer(location string) tmpl.Renderer {
	if isEnvsubstRenderer.MatchString(location) {
		return NewRenderer()
	}
	return nil
}

func init() {
	tmpl.Renderers["envsubst"] = NewRenderer
	tmpl.RendererGuesses = append(tmpl.RendererGuesses, BuildRenderer)
}

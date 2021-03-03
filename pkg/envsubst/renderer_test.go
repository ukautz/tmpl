package envsubst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRender_Render(t *testing.T) {
	vars := map[string]interface{}{
		"foo":    "foo-val",
		"BAR__1": "bar-val",
	}
	tests := []struct {
		name     string
		template string
		expect   string
		err      bool
	}{
		{`no variables used`, `This is a text`, `This is a text`, false},
		{`not defined variables untouched`, `This is a $variable unused`, `This is a $variable unused`, false},
		{`variable replacement`, `With ${foo} and ${BAR__1} variables`, `With foo-val and bar-val variables`, false},
		{`multi line replacement`, "one ${foo}\ntwo ${foo} more", "one foo-val\ntwo foo-val more", false},
	}

	renderer := NewRenderer()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rendered, err := renderer.Render(vars, []byte(test.template))
			if test.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expect, string(rendered))
			}
		})
	}
}

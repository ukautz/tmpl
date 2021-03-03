package tmpl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var _testTemplateTemplate1 = `<html>
<head><title>Status Page</title></head>
<h1>{{ .data.message }}</h1>
<ul>
	{{ range .data.announcements }}
	<li>{{ . }}</li>
	{{ end }}
</ul>
</html
`

var _testTemplateRendered1 = `<html>
<head><title>Status Page</title></head>
<h1>Things are great</h1>
<ul>
	<li>Uptime 150% in last 24h</li>
	<li>Uptime 195% in last 7 days</li>
	<li>No maintenance planned until the next century</li>
</ul>
</html
`

func TestTemplateRender_Render(t *testing.T) {
	r := TemplateRenderer{}
	res, err := r.Render(map[string]interface{}{
		"message": "Things are great",
		"announcements": []string{
			"Uptime 150% in last 24h",
			"Uptime 195% in last 7 days",
			"No maintenance planned until the next century",
		},
	}, []byte(_testTemplateTemplate1))
	assert.Nil(t, err, "No render error thrown")
	assert.Equal(t, StripTextWhitespaces(_testTemplateRendered1), StripTextWhitespaces(string(res)))
}

package pongo2

import (
	"testing"

	"github.com/stretchr/testify/assert"
	tmpl "github.com/ukautz/tmpl/pkg"
)

var _testTemplate1 = `<html>
<head><title>Status Page</title></head>
<h1>{{ data.message }}</h1>
<ul>
	{% for announcement in data.announcements %}
	<li>{{ announcement }}</li>
	{% endfor %}
</ul>
</html
`

var _testRendered1 = `<html>
<head><title>Status Page</title></head>
<h1>Things are great</h1>
<ul>
	<li>Uptime 150% in last 24h</li>
	<li>Uptime 195% in last 7 days</li>
	<li>No maintenance planned until the next century</li>
</ul>
</html
`

func TestRender_Render(t *testing.T) {
	r := &Renderer{}
	res, err := r.Render(map[string]interface{}{
		"message": "Things are great",
		"announcements": []string{
			"Uptime 150% in last 24h",
			"Uptime 195% in last 7 days",
			"No maintenance planned until the next century",
		},
	}, []byte(_testTemplate1))
	assert.Nil(t, err, "No render error thrown")
	assert.Equal(t, tmpl.StripTextWhitespaces(_testRendered1), tmpl.StripTextWhitespaces(string(res)))
}

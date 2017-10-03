package tmpl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStripTextWhitespaces(t *testing.T) {
	expects := map[string]string{
		"foo":                         "foo",
		"  foo  ":                     "foo",
		"foo\n":                       "foo",
		"foo\nbar":                    "foo\nbar",
		"  foo  \n  bar  ":            "foo\nbar",
		"  foo  \n\t\tbar \t \n baz ": "foo\nbar\nbaz",
	}
	for from, expect := range expects {
		to := StripTextWhitespaces(from)
		assert.Equal(t, expect, to, "Expected without whitespaces")
	}
}

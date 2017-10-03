package tmpl

import "strings"

func StripTextWhitespaces(from string) string {
	res := []string{}
	for _, line := range strings.Split(strings.TrimSpace(from), "\n") {
		if l := strings.TrimSpace(line); l != "" {
			res = append(res, l)
		}
	}
	return strings.Join(res, "\n")
}
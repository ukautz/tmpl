package tmpl

import "strings"

// StripTextWhitespaces returns a string which does not contain empty lines and has no whitespaces at the start
// or end of any line
func StripTextWhitespaces(from string) string {
	lines := strings.Split(strings.TrimSpace(from), "\n")
	stripped := make([]string, 0)
	for _, line := range lines {
		if line = strings.TrimSpace(line); line != "" {
			stripped = append(stripped, line)
		}
	}
	return strings.Join(stripped, "\n")
}

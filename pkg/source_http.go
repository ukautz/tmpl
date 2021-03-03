package tmpl

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

type HTTPSource string

var isHttpSource = regexp.MustCompile(`(?i)^https?://`)

// Load tries to GET given URL and return it's body content. The source must respond with HTTP status code 200
func (s HTTPSource) Load() ([]byte, error) {
	if res, err := http.Get(string(s)); err != nil {
		return nil, fmt.Errorf("could not get \"%s\": %s", s, err)
	} else if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("not ok response from \"%s\": %s", s, res.Status)
	} else if raw, err := ioutil.ReadAll(res.Body); err != nil {
		return nil, fmt.Errorf("could not read \"%s\": %s", s, err)
	} else {
		return raw, nil
	}
}

func BuildHTTPSource(location string) Source {
	if isHttpSource.MatchString(location) {
		return HTTPSource(location)
	}
	return nil
}

func init() {
	Sources = append(Sources, BuildHTTPSource)
}

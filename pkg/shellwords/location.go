package shellwords

import (
	"bytes"
	"os/exec"
	"strings"

	"github.com/mattn/go-shellwords"
	tmpl "github.com/ukautz/tmpl/pkg"
)

type Location []string

var Parser = shellwords.NewParser()

func (s Location) Load() ([]byte, error) {
	name := s[0]
	args := s[1:]
	out := bytes.NewBuffer(nil)
	cmd := exec.Command(name, args...)
	cmd.Stdout = out
	if err := cmd.Run(); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

func NewLocation(cmd string) (tmpl.Location, error) {
	parts, err := Parser.Parse(cmd)
	if err != nil {
		return nil, err
	}
	return Location(parts), nil
}

func BuildLocation(location string) tmpl.Location {
	if strings.HasPrefix(strings.ToLower(location), "shell://") {
		if src, err := NewLocation(location[len("shell://"):]); err == nil {
			return src
		}
	}
	return nil
}

func init() {
	tmpl.Locations = append(tmpl.Locations, BuildLocation)
}

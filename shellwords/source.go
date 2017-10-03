package shellwords

import (
	"bytes"
	"github.com/mattn/go-shellwords"
	"github.com/ukautz/tmpl"
	"os/exec"
	"strings"
)

type Source []string

var Parser = shellwords.NewParser()

func (s Source) Load() ([]byte, error) {
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

func NewSource(cmd string) (tmpl.Source, error) {
	parts, err := Parser.Parse(cmd)
	if err != nil {
		return nil, err
	}
	return Source(parts), nil
}

func BuildSource(location string) tmpl.Source {
	if strings.HasPrefix(strings.ToLower(location), "shell://") {
		if src, err := NewSource(location[len("shell://"):]); err == nil {
			return src
		}
	}
	return nil
}

func init() {
	tmpl.Sources = append(tmpl.Sources, BuildSource)
}

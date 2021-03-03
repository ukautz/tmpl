// Package provides an "import all" way to include all components with external dependencies with a single import statement
package imports

import (
	_ "github.com/ukautz/tmpl/pkg"
	_ "github.com/ukautz/tmpl/pkg/envsubst"
	_ "github.com/ukautz/tmpl/pkg/pongo2"
	_ "github.com/ukautz/tmpl/pkg/shellwords"
	_ "github.com/ukautz/tmpl/pkg/yaml"
)

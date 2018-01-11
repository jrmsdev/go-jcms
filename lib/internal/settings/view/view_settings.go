package view

import (
	"fmt"

	"github.com/jrmsdev/go-jcms/lib/internal/settings/args"
)

type Settings struct {
	Name     string
	Path     string
	Doctype  string
	Redirect string
	Location string
	UseView  string
	args.Args
}

func (v *Settings) ID() string {
	return fmt.Sprintf("view.%s", v.Name)
}

func (v *Settings) String() string {
	return fmt.Sprintf("%s settings", v.ID())
}

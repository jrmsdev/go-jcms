package funcs

import (
	"html/template"
	"strings"
)

var Map template.FuncMap

func init() {
	Map = template.FuncMap{
		"join": strings.Join,
	}
}

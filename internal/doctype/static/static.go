package static

import (
    "github.com/jrmsdev/go-jcms/internal/doctype"
    "github.com/jrmsdev/go-jcms/internal/doctype/base"
)

func init () {
    doctype.Register ("static", newEngine ())
}

type engine struct {
    base.Engine
}

func newEngine () *engine {
    return &engine {base.New ("static")}
}

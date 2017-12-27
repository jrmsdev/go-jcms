package text

import (
    "github.com/jrmsdev/go-jcms/internal/doctype"
    "github.com/jrmsdev/go-jcms/internal/doctype/base"
)

func init () {
    doctype.Register ("text", newEngine ())
}

type engine struct {
    base.Engine
}

func newEngine () *engine {
    return &engine {base.New ("text")}
}

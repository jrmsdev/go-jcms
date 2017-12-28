package static

import (
    "log"
    "context"
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

func (e *engine) Handle (ctx context.Context) context.Context {
    log.Println (e, "handle")
    // TODO: handle static doctype
    return ctx
}

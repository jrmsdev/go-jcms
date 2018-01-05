package text

import (
    "log"
    "context"
    "net/http"
    "github.com/jrmsdev/go-jcms/lib/internal/doctype"
    "github.com/jrmsdev/go-jcms/lib/internal/doctype/base"
    "github.com/jrmsdev/go-jcms/lib/internal/context/appctx"
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

func (e *engine) Handle (ctx context.Context) context.Context {
    log.Println (e, "handle")
    // TODO: handle text doctype
    resp := appctx.Response (ctx)
    resp.SetStatus (http.StatusOK)
    resp.Write("<html><body><p>YEAH!!!</p></body></html>")
    return ctx
}

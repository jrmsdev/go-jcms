package appctx

import (
    "context"
    "net/http"
)

type key int
const (
    ctxFail key = iota
    ctxReq
    ctxRedirect
)

func New (req *http.Request) (context.Context, context.CancelFunc) {
    ctx, cancel := context.WithCancel (context.Background ())
    ctx = context.WithValue (ctx, ctxReq, req)
    return ctx, cancel
}

func Fail (ctx context.Context) context.Context {
    return context.WithValue (ctx, ctxFail, true)
}

func Failed (ctx context.Context) bool {
    return getBool (ctx, ctxFail)
}

func Request (ctx context.Context) (*http.Request, bool) {
    req, ok := ctx.Value (ctxReq).(*http.Request)
    return req, ok
}

func SetRedirect (ctx context.Context) context.Context {
    return context.WithValue (ctx, ctxRedirect, true)
}

func Redirect (ctx context.Context) bool {
    return getBool (ctx, ctxRedirect)
}

func getBool (ctx context.Context, k key) bool {
    v, ok := ctx.Value (k).(bool)
    if !ok {
        return false
    }
    return v
}

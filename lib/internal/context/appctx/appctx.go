package appctx

import (
    "context"
    "net/http"
    "github.com/jrmsdev/go-jcms/lib/internal/response"
)

type key int
const (
    ctxFail key = iota
    ctxReq
    ctxResp
    ctxRedirect
)

func New (req *http.Request, resp *response.Response) (context.Context, context.CancelFunc) {
    ctx, cancel := context.WithCancel (context.Background ())
    ctx = context.WithValue (ctx, ctxReq, req)
    ctx = context.WithValue (ctx, ctxResp, resp)
    return ctx, cancel
}

func Fail (ctx context.Context) context.Context {
    return context.WithValue (ctx, ctxFail, true)
}

func Failed (ctx context.Context) bool {
    return getBool (ctx, ctxFail)
}

func Request (ctx context.Context) *http.Request {
    req, ok := ctx.Value (ctxReq).(*http.Request)
    if !ok {
        panic ("appctx nil request")
    }
    return req
}

func Response (ctx context.Context) *response.Response {
    req, ok := ctx.Value (ctxResp).(*response.Response)
    if !ok {
        panic ("appctx nil response")
    }
    return req
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

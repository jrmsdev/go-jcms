package appctx

import (
    "context"
    "net/http"
)

type key int
const reqKey key = 0

func New (req *http.Request) (context.Context, context.CancelFunc) {
    ctx, cancel := context.WithCancel (context.Background ())
    ctx = context.WithValue (ctx, reqKey, req)
    return ctx, cancel
}

func Request (ctx context.Context) (*http.Request, bool) {
    req, ok := ctx.Value (reqKey).(*http.Request)
    return req, ok
}

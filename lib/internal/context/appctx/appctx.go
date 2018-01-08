package appctx

import (
	"context"
	"net/http"
)

type key int

const (
	ctxFail key = iota
	ctxRedirect
)

func New(req *http.Request) (*http.Request, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	req = req.WithContext(ctx)
	return req, cancel
}

func Fail(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxFail, true)
}

func Failed(ctx context.Context) bool {
	return getBool(ctx, ctxFail)
}

func SetRedirect(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxRedirect, true)
}

func Redirect(ctx context.Context) bool {
	return getBool(ctx, ctxRedirect)
}

func getBool(ctx context.Context, k key) bool {
	v, ok := ctx.Value(k).(bool)
	if !ok {
		return false
	}
	return v
}

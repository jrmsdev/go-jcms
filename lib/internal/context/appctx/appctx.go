package appctx

import (
	"context"
)

type appkey int

const (
	ctxFail appkey = iota
	ctxRedirect
)

func New() (context.Context, context.CancelFunc) {
	return context.WithCancel(context.Background())
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

func getBool(ctx context.Context, k appkey) bool {
	v, ok := ctx.Value(k).(bool)
	if !ok {
		return false
	}
	return v
}

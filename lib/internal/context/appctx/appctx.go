package appctx

import (
	"context"
)

type appkey int

const (
	ctxFail appkey = iota
	ctxRedirect
	ctxEngineFail
)

func New() (context.Context, context.CancelFunc) {
	return context.WithCancel(context.Background())
}

func getBool(ctx context.Context, k appkey) bool {
	v, ok := ctx.Value(k).(bool)
	if !ok {
		return false
	}
	return v
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

func EngineFail(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, ctxEngineFail, true)
	return context.WithValue(ctx, ctxFail, false)
}

func EngineFailed(ctx context.Context) bool {
	return getBool(ctx, ctxEngineFail)
}

package middleware

import (
	"context"
	"net/http"

	"github.com/jrmsdev/go-jcms/lib/internal/context/appctx"
	"github.com/jrmsdev/go-jcms/lib/internal/logger"
	"github.com/jrmsdev/go-jcms/lib/internal/response"
)

var log = logger.New("middleware")
var mwreg = newRegistry()

type MiddlewareAction int

const (
	ACTION_PRE MiddlewareAction = iota
	ACTION_POST
)

type Middleware interface {
	Name() string
	Action(
		ctx context.Context,
		resp *response.Response,
		action MiddlewareAction,
		req *http.Request,
	) context.Context
}

func Register(mw Middleware, actions ...MiddlewareAction) {
	mwreg.Register(mw, actions...)
}

func Enable(settings []*Settings) error {
	return mwreg.Enable(settings)
}

func Action(
	ctx context.Context,
	resp *response.Response,
	action MiddlewareAction,
	req *http.Request,
) context.Context {
	for _, mw := range mwreg.GetAll(action) {
		if ctx := mw.Action(ctx, resp, action, req); appctx.Failed(ctx) {
			return ctx
		}
	}
	return ctx
}

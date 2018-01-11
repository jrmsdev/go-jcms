package middleware

import (
	"context"
	"net/http"

	"github.com/jrmsdev/go-jcms/lib/internal/context/appctx"
	"github.com/jrmsdev/go-jcms/lib/internal/logger"
	"github.com/jrmsdev/go-jcms/lib/internal/response"
	"github.com/jrmsdev/go-jcms/lib/internal/settings"
	"github.com/jrmsdev/go-jcms/lib/internal/settings/middleware"
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
		req *http.Request,
		cfg *settings.Reader,
		action MiddlewareAction,
	) context.Context
}

func Register(mw Middleware, actions ...MiddlewareAction) {
	mwreg.Register(mw, actions...)
}

func Enable(settings []*middleware.Settings) error {
	return mwreg.Enable(settings)
}

func Action(
	ctx context.Context,
	resp *response.Response,
	req *http.Request,
	cfg *settings.Reader,
	action MiddlewareAction,
) context.Context {
	for _, mw := range mwreg.GetAll(action) {
		cfg.SetMiddleware(mw.Name())
		ctx = mw.Action(ctx, resp, req, cfg, action)
		if appctx.Failed(ctx) {
			return ctx
		}
		cfg.Reset()
	}
	return ctx
}

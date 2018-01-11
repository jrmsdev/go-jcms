package setheader

import (
	"context"

	"github.com/jrmsdev/go-jcms/lib/internal/logger"
	"github.com/jrmsdev/go-jcms/lib/internal/middleware"
	"github.com/jrmsdev/go-jcms/lib/internal/request"
	"github.com/jrmsdev/go-jcms/lib/internal/response"
	"github.com/jrmsdev/go-jcms/lib/internal/settings"
)

const jcmsid = "middleware.setheader"

var log = logger.New(jcmsid)

func Init() {
	middleware.Register(&Middleware{}, middleware.ACTION_POST)
}

type Middleware struct{}

func (m *Middleware) Name() string {
	return "setheader"
}

func (m *Middleware) Action(
	ctx context.Context,
	resp *response.Response,
	_ *request.Request,
	cfg *settings.Reader,
	action middleware.MiddlewareAction,
) context.Context {
	if action == middleware.ACTION_PRE {
		return ctx
	}
	// global headers
	args := cfg.Middleware.Args
	for k, v := range args.GetAll("") {
		resp.SetHeader(k, v.String())
	}
	// view headers
	vargs := cfg.View.Args
	for k, v := range vargs.GetAll("") {
		resp.SetHeader(k, v.String())
	}
	return ctx
}

package setheader

import (
	"context"
	"net/http"

	"github.com/jrmsdev/go-jcms/lib/internal/logger"
	"github.com/jrmsdev/go-jcms/lib/internal/middleware"
	"github.com/jrmsdev/go-jcms/lib/internal/response"
	"github.com/jrmsdev/go-jcms/lib/internal/settings"
)

const jcmsid = "middleware.setheader"
var log = logger.New(jcmsid)

func init() {
	middleware.Register(&Middleware{}, middleware.ACTION_POST)
}

type Middleware struct{}

func (m *Middleware) Name() string {
	return "setheader"
}

func (m *Middleware) Action(
	ctx context.Context,
	resp *response.Response,
	req *http.Request,
	cfg *settings.Reader,
	action middleware.MiddlewareAction,
) context.Context {
	// TODO: setheader middleware Action
	return ctx
}

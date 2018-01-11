package setheader

import (
	"context"
	"net/http"

	"github.com/jrmsdev/go-jcms/lib/internal/logger"
	"github.com/jrmsdev/go-jcms/lib/internal/middleware"
	"github.com/jrmsdev/go-jcms/lib/internal/response"
)

var log = logger.New("middleware.setheader")

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
	action middleware.MiddlewareAction,
	req *http.Request,
) context.Context {
	return ctx
}

package templates

import (
	"context"
	"net/http"

	"github.com/jrmsdev/go-jcms/lib/internal/doctype"
	"github.com/jrmsdev/go-jcms/lib/internal/doctype/base"
	"github.com/jrmsdev/go-jcms/lib/internal/logger"
	"github.com/jrmsdev/go-jcms/lib/internal/response"
	"github.com/jrmsdev/go-jcms/lib/internal/request"
	"github.com/jrmsdev/go-jcms/lib/internal/settings"
)

var log = logger.New("doctype.templates")

func init() {
	doctype.Register("templates", newEngine())
}

type engine struct {
	base.Engine
}

func newEngine() *engine {
	return &engine{base.New("templates")}
}

func (e *engine) Handle(
	ctx context.Context,
	resp *response.Response,
	req *request.Request,
	cfg *settings.Reader,
	docroot string,
) context.Context {
	resp.SetStatus(http.StatusOK)
	return ctx
}

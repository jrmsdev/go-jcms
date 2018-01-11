package skel

import (
	"context"
	"net/http"

	"github.com/jrmsdev/go-jcms/lib/internal/doctype"
	"github.com/jrmsdev/go-jcms/lib/internal/doctype/base"
	"github.com/jrmsdev/go-jcms/lib/internal/logger"
	"github.com/jrmsdev/go-jcms/lib/internal/response"
	"github.com/jrmsdev/go-jcms/lib/internal/settings"
)

var log = logger.New("doctype.skel")

func init() {
	doctype.Register("skel", newEngine())
}

type engine struct {
	base.Engine
}

func newEngine() *engine {
	return &engine{base.New("skel")}
}

func (e *engine) Handle(
	ctx context.Context,
	resp *response.Response,
	req *http.Request,
	cfg *settings.Reader,
) context.Context {
	resp.SetStatus(http.StatusOK)
	return ctx
}

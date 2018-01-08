package static

import (
	"context"
	"log"
	"net/http"

	"github.com/jrmsdev/go-jcms/lib/internal/doctype"
	"github.com/jrmsdev/go-jcms/lib/internal/doctype/base"
	"github.com/jrmsdev/go-jcms/lib/internal/response"
)

func init() {
	doctype.Register("static", newEngine())
}

type engine struct {
	base.Engine
}

func newEngine() *engine {
	return &engine{base.New("static")}
}

func (e *engine) Handle(
	req *http.Request,
	resp *response.Response,
) context.Context {
	log.Println(e, "handle")
	ctx := req.Context()
	// TODO: handle static doctype
	return ctx
}

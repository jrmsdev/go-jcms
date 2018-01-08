package text

import (
	"context"
	"log"
	"net/http"

	"github.com/jrmsdev/go-jcms/lib/internal/doctype"
	"github.com/jrmsdev/go-jcms/lib/internal/doctype/base"
	"github.com/jrmsdev/go-jcms/lib/internal/response"
)

func init() {
	doctype.Register("text", newEngine())
}

type engine struct {
	base.Engine
}

func newEngine() *engine {
	return &engine{base.New("text")}
}

func (e *engine) Handle(
	req *http.Request,
	resp *response.Response,
) context.Context {
	log.Println(e, "handle")
	ctx := req.Context()
	// TODO: handle text doctype
	resp.SetStatus(http.StatusOK)
	resp.Write("<html><body><p>YEAH!!!</p></body></html>")
	return ctx
}

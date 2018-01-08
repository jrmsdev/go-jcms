package text

import (
	"context"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/jrmsdev/go-jcms/lib/internal/context/appctx"
	"github.com/jrmsdev/go-jcms/lib/internal/doctype"
	"github.com/jrmsdev/go-jcms/lib/internal/doctype/base"
	"github.com/jrmsdev/go-jcms/lib/internal/env"
	"github.com/jrmsdev/go-jcms/lib/internal/fsutils"
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
	// TODO: handle static doctype
	docroot := filepath.Join(env.WebappDir(), "docroot")
	if !fsutils.DirExists(docroot) {
		log.Println("E: docroot not found:", docroot)
		resp.SetError(http.StatusInternalServerError,
			"docroot not found")
		return appctx.Fail(ctx)
	}
	filename := filepath.Join(docroot, "lalala") // FIXME!!
	if !strings.HasSuffix(filename, ".txt") {
		resp.SetError(http.StatusBadRequest, "invalid request")
		return appctx.Fail(ctx)
	}
	resp.SetStatus(http.StatusOK)
	return ctx
}

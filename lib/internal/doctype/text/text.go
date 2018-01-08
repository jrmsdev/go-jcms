package text

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/jrmsdev/go-jcms/lib/internal/context/appctx"
	"github.com/jrmsdev/go-jcms/lib/internal/doctype"
	"github.com/jrmsdev/go-jcms/lib/internal/doctype/base"
	"github.com/jrmsdev/go-jcms/lib/internal/env"
	"github.com/jrmsdev/go-jcms/lib/internal/fsutils"
	"github.com/jrmsdev/go-jcms/lib/internal/response"
	"github.com/jrmsdev/go-jcms/lib/internal/views"
)

func init() {
	doctype.Register("text", newEngine())
}

const (
	maxSize = 1024 // max size to send to client
)

type engine struct {
	base.Engine
}

func newEngine() *engine {
	return &engine{base.New("text")}
}

func (e *engine) Handle(
	view *views.View,
	req *http.Request,
	resp *response.Response,
) context.Context {
	log.Println(e, "handle")
	ctx := req.Context()
	docroot := filepath.Join(env.WebappDir(), "docroot")
	if !fsutils.DirExists(docroot) {
		log.Println("E: docroot not found:", docroot)
		resp.SetError(http.StatusInternalServerError, "docroot not found")
		return appctx.Fail(ctx)
	}
	filename, ok := getFilename(view, req, docroot)
	if !ok {
		log.Println("E: file not found:", filename)
		resp.SetError(http.StatusNotFound, "file not found")
		return appctx.Fail(ctx)
	}
	err := sendFile(resp, filename)
	if err != nil {
		log.Println("E:", err)
		return appctx.Fail(ctx)
	}
	return ctx
}

func getFilename(view *views.View, req *http.Request, docroot string) (string, bool) {
	fn := req.URL.Path
	if fn == "" || fn == "/" {
		fn = path.Clean(view.Path)
	}
	filename := filepath.Join(docroot, fn+".txt")
	if !fsutils.FileExists(filename) {
		return filename, false
	}
	return filename, true
}

func sendFile(resp *response.Response, filename string) error {
	fh, err := os.Open(filename)
	if err != nil {
		resp.SetError(http.StatusInternalServerError, err.Error())
		return err
	}
	_, err = io.CopyN(resp, fh, maxSize)
	if err != nil && err != io.EOF {
		resp.SetError(http.StatusInternalServerError, err.Error())
		return err
	}
	return nil
}

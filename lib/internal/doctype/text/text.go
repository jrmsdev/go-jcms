package text

import (
	"context"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/jrmsdev/go-jcms/lib/internal/doctype"
	"github.com/jrmsdev/go-jcms/lib/internal/doctype/base"
	"github.com/jrmsdev/go-jcms/lib/internal/fsutils"
	"github.com/jrmsdev/go-jcms/lib/internal/logger"
	"github.com/jrmsdev/go-jcms/lib/internal/request"
	"github.com/jrmsdev/go-jcms/lib/internal/response"
	"github.com/jrmsdev/go-jcms/lib/internal/settings"
)

var log = logger.New("doctype.text")

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
	ctx context.Context,
	resp *response.Response,
	req *request.Request,
	cfg *settings.Reader,
	docroot string,
) context.Context {
	filename, ok := getFilename(cfg, req, docroot)
	if !ok {
		log.E("file not found:", filename)
		return resp.SetError(ctx,
			http.StatusNotFound, "file not found")
	}
	err := sendFile(resp, filename)
	if err != nil {
		log.E(err.Error())
		return resp.SetError(ctx,
			http.StatusInternalServerError, err.Error())
	}
	resp.SetStatus(http.StatusOK)
	return ctx
}

func getFilename(cfg *settings.Reader, req *request.Request, docroot string) (string, bool) {
	fn := req.URL.Path
	if fn == "" || fn == "/" {
		fn = path.Clean(cfg.View.Path)
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
		return err
	}
	defer fh.Close()
	_, err = io.CopyN(resp, fh, maxSize)
	if err != nil && err != io.EOF {
		return err
	}
	return nil
}

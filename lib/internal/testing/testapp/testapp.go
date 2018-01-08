package testapp

import (
	"context"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"testing"

	"github.com/jrmsdev/go-jcms/lib/internal/app"
	"github.com/jrmsdev/go-jcms/lib/internal/context/appctx"
	"github.com/jrmsdev/go-jcms/lib/internal/response"
)

func init() {
	setEnv("testing")
}

func setEnv(appname string) {
	os.Setenv("JCMS_WEBAPP", appname)
	os.Setenv("JCMS_BASEDIR",
		filepath.Join(os.Getenv("GOPATH"),
			"src", "github.com", "jrmsdev", "go-jcms", "webapps"))
}

func getReq(path string) *http.Request {
	req := &http.Request{}
	req.URL, _ = url.Parse("http://127.0.0.1:0" + path)
	return req
}

func reqCtx(req *http.Request) (*http.Request, context.CancelFunc) {
	return appctx.New(req)
}

type Result struct {
	Req *http.Request
	Resp *response.Response
	Ctx context.Context
	Err error
}

func Handle(t *testing.T, path string) *Result {
	var cancel context.CancelFunc
	r := &Result{}
	r.Req = getReq(path)
	r.Req, cancel = reqCtx(r.Req)
	r.Ctx = r.Req.Context()
	r.Resp = response.New()
	defer cancel()
	a, err := app.New()
	if err != nil {
		t.Fatal(err)
		r.Err = err
		return r
	}
	r.Ctx = a.Handle(r.Req, r.Resp)
	return r
}

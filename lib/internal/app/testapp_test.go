package app

import (
	"context"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/jrmsdev/go-jcms/lib/internal/context/appctx"
	"github.com/jrmsdev/go-jcms/lib/internal/response"
)

type testapp struct {
}

type testappResult struct {
	App  *App
	Req  *http.Request
	Resp *response.Response
	Ctx  context.Context
	Err  error
}

func newTestApp() *testapp {
	testappEnv("testing")
	return &testapp{}
}

func testappEnv(appname string) {
	if appname == "" {
		os.Setenv("JCMS_WEBAPP", "")
		os.Setenv("JCMS_BASEDIR", "")
		return
	}
	os.Setenv("JCMS_WEBAPP", appname)
	os.Setenv("JCMS_BASEDIR",
		filepath.Join(os.Getenv("GOPATH"),
			"src", "github.com", "jrmsdev", "go-jcms", "webapps"))
}

func (a *testapp) getReq(path string) *http.Request {
	req := &http.Request{}
	req.URL, _ = url.Parse("http://127.0.0.1:0" + path)
	return req
}

func (a *testapp) reqCtx(req *http.Request) (*http.Request, context.CancelFunc) {
	return appctx.New(req)
}

func (a *testapp) Handle(path string) *testappResult {
	var (
		err    error
		cancel context.CancelFunc
	)
	r := &testappResult{}
	r.Req = a.getReq(path)
	r.Req, cancel = a.reqCtx(r.Req)
	r.Ctx = r.Req.Context()
	r.Resp = response.New()
	defer cancel()
	r.App, err = New()
	if err != nil {
		r.Err = err
		return r
	}
	r.Ctx = r.App.Handle(r.Req, r.Resp)
	return r
}

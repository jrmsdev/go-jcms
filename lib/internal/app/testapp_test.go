package app

import (
	"context"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/jrmsdev/go-jcms/lib/internal/context/appctx"
	"github.com/jrmsdev/go-jcms/lib/internal/request"
	"github.com/jrmsdev/go-jcms/lib/internal/response"
	"github.com/jrmsdev/go-jcms/lib/internal/settings"
)

type testapp struct {
}

type testappResult struct {
	App      *App
	Req      *request.Request
	Resp     *response.Response
	Ctx      context.Context
	Err      error
	AppName  string
	Settings *settings.Settings
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
			"src", "github.com", "jrmsdev",
			"go-jcms", "webapps", "testing"))
}

func (a *testapp) getReq(ctx context.Context, path string) *request.Request {
	r := &http.Request{}
	r.URL, _ = url.Parse("http://127.0.0.1:0" + path)
	return request.New(ctx, r)
}

func (a *testapp) Handle(path string) *testappResult {
	var (
		err    error
		cancel context.CancelFunc
	)
	r := &testappResult{}
	r.AppName = "testing"
	r.Err = nil
	r.Settings, err = settings.New()
	if err != nil {
		r.Err = err
		return r
	}
	r.Ctx, cancel = appctx.New()
	defer cancel()
	r.Req = a.getReq(r.Ctx, path)
	r.Resp = response.New()
	r.App, err = New(r.AppName, r.Settings)
	if err != nil {
		r.Err = err
		return r
	}
	r.Ctx = r.App.Handle(r.Ctx, r.Resp, r.Req)
	return r
}

package app

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"testing"

	"github.com/jrmsdev/go-jcms/lib/internal/context/appctx"
	"github.com/jrmsdev/go-jcms/lib/internal/response"
)

func init() {
	setEnv("testing")
}

func TestFindView(t *testing.T) {
	req := getReq("/test")
	a, err := New()
	if err != nil {
		t.Fatal(err)
	}
	v, err := a.findView(req.URL.Path)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(v)
}

func TestViewNotFound(t *testing.T) {
	req := getReq("/notfound")
	a, err := New()
	if err != nil {
		t.Fatal(err)
	}
	v, err := a.findView(req.URL.Path)
	if err == nil {
		t.Fatal("found view:", v.Name)
	}
}

func TestNewApp(t *testing.T) {
	a, err := New()
	if err != nil {
		t.Fatal(err)
	}
	if a.String() != fmt.Sprintf("app.%s", a.name) {
		t.Error("a.String != app.<name>")
	}
}

func TestNewAppSettingsError(t *testing.T) {
	setEnv("invalidapp")
	a, err := New()
	if err == nil {
		t.Log(a, err)
		t.Error("settings file for invalidapp should fail")
	}
	setEnv("testing") // restore env
}

func setEnv(appname string) {
	os.Setenv("JCMS_WEBAPP", appname)
	os.Setenv("JCMS_BASEDIR",
		filepath.Join(os.Getenv("GOPATH"),
			"src", "github.com", "jrmsdev", "go-jcms", "apps"))
}

func getReq(path string) *http.Request {
	req := &http.Request{}
	req.URL, _ = url.Parse("http://127.0.0.1:0" + path)
	return req
}

func reqCtx(req *http.Request) (*http.Request, context.CancelFunc) {
	return appctx.New(req)
}

func TestAppHandle(t *testing.T) {
	req := getReq("/test")
	req, cancel := reqCtx(req)
	ctx := req.Context()
	resp := response.New()
	defer cancel()
	a, err := New()
	if err != nil {
		t.Fatal(err)
	}
	ctx = a.Handle(req, resp)
	if appctx.Failed(ctx) {
		t.Error("app.Handle should not fail")
	}
}

func TestAppHandleViewNotFound(t *testing.T) {
	req := getReq("/test/view.not.found")
	req, cancel := reqCtx(req)
	ctx := req.Context()
	resp := response.New()
	defer cancel()
	a, err := New()
	if err != nil {
		t.Fatal(err)
	}
	ctx = a.Handle(req, resp)
	if !appctx.Failed(ctx) {
		t.Fatal("app.Handle should fail")
	}
	if resp.Error() != "view: not found /test/view.not.found" {
		t.Log(resp.Error())
		t.Error("wrong app.Handle view not found error message")
	}
}

func TestAppHandleInvalidEngine(t *testing.T) {
	req := getReq("/test/doctype.engine.invalid")
	req, cancel := reqCtx(req)
	ctx := req.Context()
	resp := response.New()
	defer cancel()
	a, err := New()
	if err != nil {
		t.Fatal(err)
	}
	ctx = a.Handle(req, resp)
	if !appctx.Failed(ctx) {
		t.Fatal("app.Handle should fail")
	}
	if resp.Error() != "invalid doctype engine: invalid.engine" {
		t.Log(resp.Error())
		t.Error("wrong app.Handle invalid doctype engine error message")
	}
}

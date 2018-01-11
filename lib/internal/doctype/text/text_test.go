package text

import (
	"context"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jrmsdev/go-jcms/lib/internal/context/appctx"
	"github.com/jrmsdev/go-jcms/lib/internal/doctype"
	"github.com/jrmsdev/go-jcms/lib/internal/env"
	"github.com/jrmsdev/go-jcms/lib/internal/response"
	"github.com/jrmsdev/go-jcms/lib/internal/settings"
	"github.com/jrmsdev/go-jcms/lib/internal/settings/view"
)

func testappEnv(appname string) {
	if appname == "" {
		os.Setenv("JCMS_WEBAPP", "")
		os.Setenv("JCMS_BASEDIR", "")
	}
	os.Setenv("JCMS_WEBAPP", appname)
	os.Setenv("JCMS_BASEDIR",
		filepath.Join(os.Getenv("GOPATH"),
			"src", "github.com", "jrmsdev",
			"go-jcms", "webapps", "testing"))
}

func getReq(ctx context.Context, path string) *http.Request {
	r := &http.Request{}
	req := r.WithContext(ctx)
	req.URL, _ = url.Parse("http://127.0.0.1:0" + path)
	return req
}

func getCfg(
	name string,
	path string,
	dtype string,
) (*settings.Reader, error) {
	s, err := settings.New(env.SettingsFile())
	if err != nil {
		return nil, err
	}
	v := &view.Settings{Name: name, Path: path, Doctype: dtype}
	return settings.NewReader(s, v), nil
}

func TestEngine(t *testing.T) {
	e, err := doctype.GetEngine("text")
	if err != nil {
		t.Fatal(err)
	}
	testType(t, e)
	testHandle(t, e)
	testHandleDocrootError(t, e)
	testHandleNotFound(t, e)
}

func testType(t *testing.T, e doctype.Engine) {
	if e.Type() != "text" {
		t.Error(".Type != text")
	}
}

func testHandle(t *testing.T, e doctype.Engine) {
	testappEnv("testing")
	defer testappEnv("")
	cfg, err := getCfg("testview", "/pathto/testview", "text")
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := appctx.New()
	defer cancel()
	req := getReq(ctx, "/test")
	resp := response.New()
	ctx = e.Handle(ctx, resp, req, cfg)
	if appctx.Failed(ctx) {
		t.Error("handle context should not fail:", resp.Error())
	}
	status := resp.Status()
	if status != http.StatusOK {
		t.Error("invalid resp status:", status)
	}
	body := strings.TrimSpace(string(resp.Body()))
	if body != "testing" {
		t.Error("invalid resp body:", body)
	}
}

func testHandleDocrootError(t *testing.T, e doctype.Engine) {
	testappEnv("nodocroot")
	defer testappEnv("")
	cfg, err := getCfg("testview", "/pathto/testview", "text")
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := appctx.New()
	defer cancel()
	req := getReq(ctx, "/")
	resp := response.New()
	ctx = e.Handle(ctx, resp, req, cfg)
	if !appctx.Failed(ctx) {
		t.Error("handle context has not failed")
	}
	errmsg := resp.Error()
	if errmsg != "docroot not found" {
		t.Error("invalid resp error:", errmsg)
	}
}

func testHandleNotFound(t *testing.T, e doctype.Engine) {
	testappEnv("testing")
	defer testappEnv("")
	cfg, err := getCfg("testview", "/pathto/testview", "text")
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := appctx.New()
	defer cancel()
	req := getReq(ctx, "/invaliduri")
	resp := response.New()
	ctx = e.Handle(ctx, resp, req, cfg)
	if !appctx.Failed(ctx) {
		t.Error("handle context should fail")
	}
	status := resp.Status()
	if status != http.StatusNotFound {
		t.Error("invalid resp status:", status)
	}
	errmsg := resp.Error()
	if errmsg != "file not found" {
		t.Error("invalid error message:", errmsg)
	}
}

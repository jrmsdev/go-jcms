package skel

import (
	"context"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
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
	e, err := doctype.GetEngine("skel")
	if err != nil {
		t.Fatal(err)
	}
	testType(t, e)
	testHandle(t, e)
}

func testType(t *testing.T, e doctype.Engine) {
	if e.Type() != "skel" {
		t.Error(".Type != skel")
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
}

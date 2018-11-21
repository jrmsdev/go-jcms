package testeng

import (
	"context"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"testing"

	"github.com/jrmsdev/go-jcms/lib/internal/context/appctx"
	"github.com/jrmsdev/go-jcms/lib/internal/doctype"
	"github.com/jrmsdev/go-jcms/lib/internal/request"
	"github.com/jrmsdev/go-jcms/lib/internal/response"
	"github.com/jrmsdev/go-jcms/lib/internal/settings"
	"github.com/jrmsdev/go-jcms/lib/internal/settings/view"
)

type Result struct {
	Ctx     context.Context
	Resp    *response.Response
	Req     *request.Request
	Cfg     *settings.Reader
	Docroot string
}

func (r *Result) Cleanup() {
	r.Ctx = nil
	r.Resp = nil
	r.Req = nil
	r.Cfg = nil
	r.Docroot = ""
	testappEnv("")
}

type Query struct {
	App     string
	Name    string
	Path    string
	Doctype string
}

func (q *Query) setDefaults() {
	if q.App == "" {
		q.App = "testing"
	}
	if q.Name == "" {
		q.Name = "test"
	}
	if q.Path == "" {
		q.Path = "/test"
	}
	if q.Doctype == "" {
		q.Doctype = "text"
	}
}

func Handle(t *testing.T, name string, q *Query) *Result {
	var (
		err    error
		cancel context.CancelFunc
	)
	q.setDefaults()
	eng, err := doctype.GetEngine(name)
	if err != nil {
		t.Fatal(err)
	}
	r := &Result{}
	r.Docroot = testappEnv(q.App)
	r.Cfg, err = getCfg(q.Name, q.Path, q.Doctype)
	if err != nil {
		t.Fatal(err)
	}
	r.Ctx, cancel = appctx.New()
	defer cancel()
	r.Req = getReq(r.Ctx, q.Path)
	r.Resp = response.New()
	r.Ctx = eng.Handle(r.Ctx, r.Resp, r.Req, r.Cfg, r.Docroot)
	return r
}

func testappEnv(appname string) string {
	if appname == "" {
		os.Setenv("JCMS_WEBAPP", "")
		os.Setenv("JCMS_BASEDIR", "")
		return ""
	}
	basedir := filepath.Join(os.Getenv("GOPATH"),
		"src", "github.com", "jrmsdev",
		"go-jcms", "webapps", "testing")
	os.Setenv("JCMS_WEBAPP", appname)
	os.Setenv("JCMS_BASEDIR", basedir)
	return filepath.Join(basedir, appname, "docroot")
}

func getCfg(
	name string,
	path string,
	dtype string,
) (*settings.Reader, error) {
	s, err := settings.New()
	if err != nil {
		return nil, err
	}
	v := &view.Settings{Name: name, Path: path, Doctype: dtype}
	return settings.NewReader(s, v), nil
}

func getReq(ctx context.Context, path string) *request.Request {
	r := &http.Request{}
	r.URL, _ = url.Parse("http://127.0.0.1:0" + path)
	return request.New(ctx, r)
}

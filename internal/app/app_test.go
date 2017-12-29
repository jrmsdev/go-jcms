package app

import (
    "os"
    "fmt"
    "testing"
    "context"
    "net/url"
    "net/http"
    "path/filepath"
    "github.com/jrmsdev/go-jcms/internal/response"
    "github.com/jrmsdev/go-jcms/internal/context/appctx"
)

func init () {
    setEnv ("devel")
}

func TestFindView (t *testing.T) {
    req := getReq ("/test")
    a, err := New ()
    if err != nil {
        t.Fatal (err)
    }
    v, err := a.findView (req.URL.Path)
    if err != nil {
        t.Fatal (err)
    }
    t.Log (v)
}

func TestViewNotFound (t *testing.T) {
    req := getReq ("/notfound")
    a, err := New ()
    if err != nil {
        t.Fatal (err)
    }
    v, err := a.findView (req.URL.Path)
    if err == nil {
        t.Fatal ("found view:", v.Name)
    }
}

func TestNewApp (t *testing.T) {
    a, err := New ()
    if err != nil {
        t.Fatal (err)
    }
    if a.String () != fmt.Sprintf ("app.%s", a.name) {
        t.Error ("a.String != app.<name>")
    }
}

func TestNewAppSettingsError (t *testing.T) {
    setEnv ("invalidapp")
    a, err := New ()
    if err == nil {
        t.Log (a, err)
        t.Error ("settings file for invalidapp should fail")
    }
    setEnv ("devel") // restore env
}

func setEnv (appname string) {
    os.Setenv ("JCMS_WEBAPP", appname)
    os.Setenv ("JCMS_BASEDIR",
            filepath.Join (os.Getenv ("GOPATH"),
                    "src", "github.com", "jrmsdev", "go-jcms", "apps"))
}

func getReq (path string) *http.Request {
    req := &http.Request{}
    req.URL, _ = url.Parse ("http://127.0.0.1:0" + path)
    return req
}

func getCtx (path string) (context.Context, context.CancelFunc) {
    return appctx.New (getReq (path), response.New ())
}

func TestAppHandle (t *testing.T) {
    ctx, cancel := getCtx ("/test")
    defer cancel()
    a, err := New ()
    if err != nil {
        t.Fatal (err)
    }
    ctx = a.Handle (ctx)
    if appctx.Failed (ctx) {
        t.Error ("app.Handle should not fail")
    }
}

func TestAppHandleViewNotFound (t *testing.T) {
    ctx, cancel := getCtx ("/test/view.not.found")
    defer cancel()
    a, err := New ()
    if err != nil {
        t.Fatal (err)
    }
    ctx = a.Handle (ctx)
    if !appctx.Failed (ctx) {
        t.Fatal ("app.Handle should fail")
    }
    resp := appctx.Response (ctx)
    if resp.Error () != "view: not found /test/view.not.found" {
        t.Log (resp.Error ())
        t.Error ("wrong app.Handle view not found error message")
    }
}

func TestAppHandleInvalidEngine (t *testing.T) {
    ctx, cancel := getCtx ("/test/doctype.engine.invalid")
    defer cancel()
    a, err := New ()
    if err != nil {
        t.Fatal (err)
    }
    ctx = a.Handle (ctx)
    if !appctx.Failed (ctx) {
        t.Fatal ("app.Handle should fail")
    }
    resp := appctx.Response (ctx)
    if resp.Error () != "invalid doctype engine: invalid.engine" {
        t.Log (resp.Error ())
        t.Error ("wrong app.Handle invalid doctype engine error message")
    }
}

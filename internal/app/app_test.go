package app

import (
    "os"
    "fmt"
    "testing"
    //~ "context"
    "net/url"
    "net/http"
    "path/filepath"
    //~ "github.com/jrmsdev/go-jcms/internal/context/appctx"
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

//~ func getCtx (req *http.Request) (context.Context, context.CancelFunc) {
    //~ return appctx.New (req)
//~ }

func getReq (path string) *http.Request {
    req := &http.Request{}
    req.URL, _ = url.Parse ("http://127.0.0.1:0" + path)
    return req
}

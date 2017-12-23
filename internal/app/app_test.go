package app

import (
    "os"
    "testing"
    //~ "context"
    "net/url"
    "net/http"
    "path/filepath"
    //~ "github.com/jrmsdev/go-jcms/internal/context/appctx"
)

func init () {
    os.Setenv ("JCMS_WEBAPP", "devel")
    os.Setenv ("JCMS_BASEDIR",
            filepath.Join (os.Getenv ("GOPATH"),
                    "src", "github.com", "jrmsdev", "go-jcms", "apps"))
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

//~ func getCtx (req *http.Request) (context.Context, context.CancelFunc) {
    //~ return appctx.New (req)
//~ }

func getReq (path string) *http.Request {
    req := &http.Request{}
    req.URL, _ = url.Parse ("http://127.0.0.1:0" + path)
    return req
}

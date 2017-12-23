package app

import (
    "os"
    "testing"
    "context"
    "net/http"
    "path/filepath"
    "github.com/jrmsdev/go-jcms/internal/context/appctx"
)

func init () {
    os.Setenv ("JCMS_WEBAPP", "devel")
    os.Setenv ("JCMS_BASEDIR",
            filepath.Join (os.Getenv ("GOPATH"),
                    "src", "github.com", "jrmsdev", "go-jcms", "apps"))
}

func TestFindView (t *testing.T) {
    //~ ctx, cancel := getCtx ()
    //~ defer cancel()
    a, err := New ()
    if err != nil {
        t.Fatal (err)
    }
    a.findView ()
}

func getCtx () (context.Context, context.CancelFunc) {
    req := &http.Request{}
    return appctx.New (req)
}

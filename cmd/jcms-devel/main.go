package main

import (
    "os"
    "flag"
    "path"
    "github.com/jrmsdev/go-jcms/internal/cli/webview"
)

func main () {
    // use devel webapp
    if err := os.Setenv ("JCMS_WEBAPP", "devel"); err != nil {
        panic (err)
    }
    // force using GOPATH
    if err := os.Setenv ("JCMS_BASEDIR", ""); err != nil {
        panic (err)
    }
    flag.Parse ()
    req := path.Clean (flag.Arg (0))
    if req == "." || req == "" {
        req = "/"
    }
    webview.DevelMain (req)
}

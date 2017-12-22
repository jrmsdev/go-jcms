package main

import (
    "os"
    "github.com/jrmsdev/go-jcms/jcms/cli"
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
    cli.Webview ()
}

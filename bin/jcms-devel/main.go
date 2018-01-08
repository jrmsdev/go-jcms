package main

import (
    "os"
    "log"
    "flag"
    "path"
    "path/filepath"
)

func main () {
    // use devel webapp
    if err := os.Setenv ("JCMS_WEBAPP", "devel"); err != nil {
        log.Fatalln (err)
    }
    // force using GOPATH
    gp, ok := os.LookupEnv ("GOPATH")
    gp = filepath.Clean (gp)
    if !ok || gp == "." {
        log.Fatalln ("GOPATH is not set")
    }
    // TODO: support possible ':' separator in GOPATH
    gp = filepath.Join (gp, "src", "github.com", "jrmsdev", "go-jcms", "apps")
    if err := os.Setenv ("JCMS_BASEDIR", gp); err != nil {
        log.Fatalln (err)
    }
    // parse command args
    flag.Parse ()
    uri := path.Clean (flag.Arg (0))
    if uri == "." || uri == "" {
        uri = "/"
    }
    // launch webview
    Webview (uri)
}

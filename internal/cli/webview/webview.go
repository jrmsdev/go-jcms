package webview

import (
    "log"
    "net/url"
    xwv "github.com/zserge/webview"
    "github.com/jrmsdev/go-jcms/jcms"
    "github.com/jrmsdev/go-jcms/internal/rt"
)

const (
    webviewResize = true
    webviewWidth = 800
    webviewHeight = 600
)

func Main () {
    doMain ("/")
}

func DevelMain (req string) {
    doMain (req)
}

func doMain (req string) {
    uri, err := url.Parse (jcms.Listen ())
    if err != nil {
        panic (err)
    }
    log.Println ("webview: req", req)
    go func() {
        jcms.Serve ()
    }()
    uri.Path = req
    log.Println ("webview: open", uri.String ())
    xwv.Open (rt.NAME, uri.String (), webviewWidth, webviewHeight, webviewResize)
    jcms.Stop ()
}

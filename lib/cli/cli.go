// public cmd API
package cli

import (
    "log"
    "net/url"
    xwv "github.com/zserge/webview"
    "github.com/jrmsdev/go-jcms/lib/internal/rt"
    "github.com/jrmsdev/go-jcms/lib/internal/core"
)

const (
    webviewResize = true
    webviewWidth = 800
    webviewHeight = 600
)

func Main () {
    core.Listen ()
    core.Serve ()
}

func Webview (req string) {
    uri, err := url.Parse (core.Listen ())
    if err != nil {
        log.Fatalln (err)
    }
    log.Println ("webview: req", req)
    go func() {
        core.Serve ()
    }()
    uri.Path = req
    log.Println ("webview: open", uri.String ())
    xwv.Open (rt.NAME, uri.String (), webviewWidth, webviewHeight, webviewResize)
    core.Stop ()
}

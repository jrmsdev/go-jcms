package cli

import (
    "log"
    "github.com/zserge/webview"
    "github.com/jrmsdev/go-jcms/jcms"
    "github.com/jrmsdev/go-jcms/internal/rt"
)

const (
    webviewResize = true
    webviewWidth = 800
    webviewHeight = 600
)

func Webview () {
    uri := jcms.Listen ()
    go func() {
        jcms.Serve ()
    }()
    log.Println ("webview: open")
    webview.Open (rt.NAME, uri, webviewWidth, webviewHeight, webviewResize)
    jcms.Stop ()
}

package webview

import (
    "log"
    "net/url"
    "strings"
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
    go func() {
        jcms.Serve ()
    }()
    log.Println ("webview: req", req)
    if req != "/" {
        if strings.HasPrefix (req, "/") {
            req := strings.Replace (req, "/", "", 1)
            log.Println ("webview: removed req / prefix:", req)
        }
        uri.Path = req
        log.Println ("webview: final uri:", uri.String ())
    }
    log.Println ("webview: open", uri.String ())
    xwv.Open (rt.NAME, uri.String (), webviewWidth, webviewHeight, webviewResize)
    jcms.Stop ()
}

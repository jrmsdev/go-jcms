package cli

import (
    "github.com/zserge/webview"
    "github.com/jrmsdev/go-jcms/jcms"
)

func Webview () {
    uri := jcms.Listen ()

    go func() {
        jcms.Serve ()
    }()
    println (uri)

    resize := true
    webview.Open ("jcms", uri, 800, 600, resize)

    jcms.Stop ()
}

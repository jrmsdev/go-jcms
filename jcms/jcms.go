// public API
package jcms

import (
    "github.com/jrmsdev/go-jcms/internal/httpd"
    "github.com/jrmsdev/go-jcms/internal/webapps"
)

func Listen () string {
    return httpd.Listen ()
}

func Serve () {
    webapps.Start ()
    httpd.Serve ()
}

func Stop () {
    httpd.Stop ()
}

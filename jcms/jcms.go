// public API
package jcms

import (
    "github.com/jrmsdev/go-jcms/internal/httpd"
    _ "github.com/jrmsdev/go-jcms/internal/webapps"
)

func Listen () string {
    return httpd.Listen ()
}

func Serve () {
    httpd.Serve ()
}

func Stop () {
    httpd.Stop ()
}

// public API
package jcms

import (
    "log"
    "github.com/jrmsdev/go-jcms/internal/httpd"
    "github.com/jrmsdev/go-jcms/internal/config"
    "github.com/jrmsdev/go-jcms/internal/webapps"
)

var listening = false

func Listen () string {
    log.Printf ("%s version %s\n", config.NAME, Version ())
    uri := httpd.Listen ()
    log.Println ("URI:", uri)
    listening = true
    return uri
}

func Serve () {
    if listening {
        webapps.Start ()
        httpd.Serve ()
    } else {
        log.Fatalln ("E: call jcms.Listen() first")
    }
}

func Stop () {
    if listening {
        httpd.Stop ()
        listening = false
    } else {
        log.Fatalln ("E: trying to stop a not listening server...")
        log.Fatalln ("E: jcms.Listen() and jcms.Server() should be called first")
    }
}

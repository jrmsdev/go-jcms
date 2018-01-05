package core

import (
    "log"
    "github.com/jrmsdev/go-jcms/lib/jcms/version"
    "github.com/jrmsdev/go-jcms/lib/internal/rt"
    "github.com/jrmsdev/go-jcms/lib/internal/httpd"
    "github.com/jrmsdev/go-jcms/lib/internal/webapps"
)

var listening = false
var webappsStarted = false

func Listen () string {
    log.Printf ("%s version %s\n", rt.NAME, version.String ())
    uri := httpd.Listen ()
    log.Println ("URI:", uri.String ())
    listening = true
    return uri.String ()
}

func Serve () {
    if listening {
        if !webappsStarted {
            webapps.Start ()
            webappsStarted = true
        }
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
        log.Println ("E: trying to stop a not listening server...")
        log.Fatalln ("E: jcms.Listen() and jcms.Server() should be called first")
    }
}

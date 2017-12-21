package webapps

import (
    "log"
    "net/http"
    "github.com/jrmsdev/go-jcms/internal/httpd"
    "github.com/jrmsdev/go-jcms/internal/config"
)

// defaults for OS env overwritable settings
var (
    datadir = "/var/opt/jcms/data" // FIXME
    webapp = "default"
)

func setHandler () {
    httpd.HandleFunc ("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("<html><body><p>welcome to jcms!</p></body></html>"))
    })
}

func Start () {
    log.Println ("datadir:", config.Datadir ())
    log.Println ("webapp:", config.Webapp ())
    setHandler ()
}

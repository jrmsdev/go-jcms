package webapps

import (
    "log"
    "errors"
    "net/http"
    "github.com/jrmsdev/go-jcms/internal/httpd"
    "github.com/jrmsdev/go-jcms/internal/config"
    "github.com/jrmsdev/go-jcms/internal/utils"
)

func mainHandler () {
    httpd.HandleFunc ("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("<html><body><p>welcome to jcms!</p></body></html>"))
    })
}

func errHandler (msg string) {
    httpd.HandleFunc ("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("<html><body><h2>ERROR: "))
        w.Write([]byte(msg))
        w.Write([]byte("</h2></body></html>"))
    })
}

func loadWebapp () error {
    log.Println ("webapp:", config.Webapp ())
    wadir := config.WebappDir ()
    log.Println ("webapp:", wadir)
    if !utils.DirExists (wadir) {
        log.Println ("E: webapp:", wadir, "dir not found")
        return errors.New ("dir not found")
    }
    return nil
}

func Start () {
    log.Println ("webapps: start")
    var err error
    if err = loadWebapp (); err != nil {
        errHandler ("webapp: " + err.Error ())
        return
    }
    mainHandler ()
}

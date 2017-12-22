package webapps

import (
    "log"
    "html"
    "net/http"
    "github.com/jrmsdev/go-jcms/internal/app"
    "github.com/jrmsdev/go-jcms/internal/httpd"
)

func mainHandler () {
    httpd.HandleFunc ("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("<html><body><p>welcome to jcms!</p></body></html>"))
    })
}

func errHandler (msg string) {
    log.Println ("ERROR:", msg)
    httpd.HandleFunc ("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("<html><body><h2>ERROR: "))
        w.Write([]byte(html.EscapeString (msg)))
        w.Write([]byte("</h2></body></html>"))
    })
}

func Start () {
    log.Println ("webapps: start")
    var err error
    // new app
    if _, err = app.New (); err != nil {
        errHandler (err.Error ())
        return
    }
    mainHandler ()
}

package webapps

import (
    "log"
    "net/http"
    "github.com/jrmsdev/go-jcms/internal/app"
    "github.com/jrmsdev/go-jcms/internal/httpd"
)

func Start () {
    log.Println ("webapps: start")
    a, err := app.New ()
    if err != nil {
        errHandler (err)
        return
    }
    mainHandler (a)
}

func mainHandler (a *app.App) {
    log.Println ("main handler:", a)
    httpd.HandleFunc ("/", func(w http.ResponseWriter, r *http.Request) {
        resp := a.Handle (r)
        if resp.IsError () {
            respError (w, resp.Error ())
        } else {
            writeResp (w, resp)
        }
    })
}

func errHandler (err error) {
    httpd.HandleFunc ("/", func(w http.ResponseWriter, r *http.Request) {
        log.Println ("INTERNAL ERROR:", err.Error ())
        http.Error (w, "INTERNAL ERROR: " + err.Error (),
                http.StatusInternalServerError)
    })
}

func respError (w http.ResponseWriter, err *app.Error) {
    log.Println ("ERROR:", err.Error ())
    http.Error (w, "ERROR: " + err.Error (), err.Status ())
}

func writeResp (w http.ResponseWriter, resp *app.Response) {
    log.Println ("write response")
    _, err := w.Write (resp.Body ())
    if err != nil {
        log.Fatalln ("PANIC: write response:", err.Error ())
    }
}

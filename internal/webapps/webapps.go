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
    httpd.HandleFunc ("/", func(w http.ResponseWriter, r *http.Request) {
        log.Println ("main handler:", a)
        resp := a.Handle (r)
        if resp.IsError () {
            respError (w, resp.Error ())
        } else {
            writeResp (w, resp)
        }
    })
}

func errHandler (err error) {
    log.Println ("INTERNAL ERROR:", err.Error ())
    httpd.HandleFunc ("/", func(w http.ResponseWriter, r *http.Request) {
        http.Error (w, "INTERNAL ERROR: " + err.Error (),
                http.StatusInternalServerError)
    })
}

func respError (w http.ResponseWriter, err *app.Error) {
    http.Error (w, "ERROR: " + err.Error (), err.Status ())
}

func writeResp (w http.ResponseWriter, resp *app.Response) {
    log.Println ("write response")
    _, err := w.Write (resp.Body ())
    if err != nil {
        log.Println ("WARNING: ignored error writing response -",
                err.Error ())
    }
}

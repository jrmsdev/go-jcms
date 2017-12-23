package app

import (
    "log"
    "fmt"
    "errors"
    "net/http"
    "github.com/jrmsdev/go-jcms/internal/rt"
    "github.com/jrmsdev/go-jcms/internal/utils"
    "github.com/jrmsdev/go-jcms/internal/views"
)

type App struct {
    name string
    settings *Settings
    views *views.Registry
}

func New () (*App, error) {
    name := rt.WebappName ()
    log.Println ("app:", name)
    s, err := getSettings ()
    if err != nil {
        return nil, err
    }
    reg := views.Register (s.Views)
    a := &App{name, s, reg}
    return a, nil
}

func getSettings () (*Settings, error) {
    fn := rt.SettingsFile ()
    log.Println ("app:", fn)
    if !utils.FileExists (fn) {
        return nil, errors.New ("file not found: " + fn)
    }
    return readSettings (fn)
}

func (a *App) String () string {
    return fmt.Sprintf ("<app:%s>", a.name)
}

func (a *App) Handle (req *http.Request) *Response {
    // TODO: ...
    resp := newResponse ()
    resp.Write("<html><body><p>YEAH!!!</p></body></html>")
    return resp
}

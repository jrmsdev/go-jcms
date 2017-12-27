package app

import (
    "log"
    "fmt"
    "errors"
    "context"
    "github.com/jrmsdev/go-jcms/internal/rt"
    "github.com/jrmsdev/go-jcms/internal/utils"
    "github.com/jrmsdev/go-jcms/internal/views"
    "github.com/jrmsdev/go-jcms/internal/doctype"
    "github.com/jrmsdev/go-jcms/internal/context/appctx"
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

func (a *App) String () string {
    return fmt.Sprintf ("<app:%s>", a.name)
}

func (a *App) Handle (ctx context.Context) (*Response, context.Context) {
    var (
        err error
        view *views.View
        eng doctype.Engine
    )
    req := appctx.Request (ctx)
    view, err = a.findView (req.URL.Path)
    resp := newResponse ()
    if err != nil {
        resp.SetError (500, err.Error ())
        ctx = appctx.Fail (ctx)
        return resp, ctx
    }
    log.Println ("app: view doctype", view.Doctype)
    // TODO: check view doctype and handle request
    eng, err = doctype.GetEngine (view.Doctype)
    if err != nil {
        resp.SetError (500, err.Error ())
        ctx = appctx.Fail (ctx)
        return resp, ctx
    }
    log.Println ("app: view engine", eng.String ())
    resp.Write("<html><body><p>YEAH!!!</p></body></html>")
    return resp, ctx
}

func (a *App) findView (path string) (*views.View, error) {
    return a.views.Get (path)
}

func getSettings () (*Settings, error) {
    fn := rt.SettingsFile ()
    log.Println ("app:", fn)
    if !utils.FileExists (fn) {
        return nil, errors.New ("file not found: " + fn)
    }
    return readSettings (fn)
}

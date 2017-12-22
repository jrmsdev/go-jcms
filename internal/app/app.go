package app

import (
    "log"
    "fmt"
    "errors"
    "github.com/jrmsdev/go-jcms/internal/rt"
    "github.com/jrmsdev/go-jcms/internal/utils"
)

type App struct {
    name string
    cfg *Settings
}

func (a *App) String () string {
    return fmt.Sprintf ("<app:%s>", a.name)
}

func appLoad () (*Settings, error) {
    fn := rt.SettingsFile ()
    log.Println ("app:", fn)
    if !utils.FileExists (fn) {
        return nil, errors.New ("file not found: " + fn)
    }
    return readSettings (fn)
}

func New () (*App, error) {
    name := rt.WebappName ()
    log.Println ("app:", name)
    cfg, err := appLoad ()
    if err != nil {
        return nil, err
    }
    return &App{name, cfg}, nil
}

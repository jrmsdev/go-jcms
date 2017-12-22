package app

import (
    "log"
    "errors"
    "github.com/jrmsdev/go-jcms/internal/utils"
    "github.com/jrmsdev/go-jcms/internal/config"
)

type App struct {
    name string
    cfg *Settings
}

func appLoad () (*Settings, error) {
    fn := config.SettingsFile ()
    log.Println ("app:", fn)
    if !utils.FileExists (fn) {
        return nil, errors.New ("file not found: " + fn)
    }
    return readSettings (fn)
}

func New () (*App, error) {
    name := config.WebappName ()
    log.Println ("app:", name)
    cfg, err := appLoad ()
    if err != nil {
        return nil, err
    }
    return &App{name, cfg}, nil
}

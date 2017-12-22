package app

import (
    "log"
    "errors"
    "github.com/jrmsdev/go-jcms/internal/utils"
    "github.com/jrmsdev/go-jcms/internal/config"
)

type App struct {
    name string
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
    log.Println ("app:", config.WebappName ())
    _, err := appLoad ()
    if err != nil {
        return nil, err
    }
    return nil, nil
}

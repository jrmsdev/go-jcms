package config

import (
    "os"
    "path/filepath"
)

const (
    NAME = "jcms"
)

// OS env defaults
var (
    webapp = "default"
    datadir = filepath.FromSlash ("/opt/jcms")
)

func getEnv (n, d string) string {
    v, isSet := os.LookupEnv (n)
    if isSet {
        return v
    }
    return d
}

func Webapp () string {
    return getEnv ("JCMS_WEBAPP", webapp)
}

func Datadir () string {
    return getEnv ("JCMS_DATADIR", datadir)
}

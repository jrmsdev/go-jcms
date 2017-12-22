package config

import (
    "os"
    "log"
    fp "path/filepath"
)

const (
    NAME = "jcms"
)

var (
    // OS env defaults
    webapp = "devel"
    datadir = fp.FromSlash ("/opt/jcms")
)

func getEnv (n, d string) string {
    v, isSet := os.LookupEnv (n)
    if isSet {
        return v
    }
    return d
}

func absPath (p string) string {
    rp, err := fp.Abs (fp.Clean (p))
    if err != nil {
        log.Fatalln ("E: config absPath:", p, "-", err)
    }
    return rp
}

func baseDatadir () string {
    return absPath (getEnv ("JCMS_DATADIR", datadir))
}

func WebappName () string {
    return getEnv ("JCMS_WEBAPP", webapp)
}

func Datadir () string {
    return absPath (fp.Join (baseDatadir(), WebappName ()))
}

func SettingsFile () string {
    return absPath (fp.Join (Datadir (), "webapp.xml"))
}

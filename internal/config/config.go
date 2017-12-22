package config

import (
    "os"
    "log"
    fp "path/filepath"
)

const NAME = "jcms"

// OS env defaults - set as JCMS_<UPPERCASE_NAME> - ie: JCMS_WEBAPP
// JCMS_BASEDIR env var is dinamically checked
var webapp = "default"

func WebappName () string {
    return getEnv ("JCMS_WEBAPP", webapp)
}

func SettingsFile () string {
    return absPath (fp.Join (webappDir (), "webapp.xml"))
}

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

func baseDir () string {
    v, ok := os.LookupEnv ("JCMS_BASEDIR")
    if ok && v != "" {
        return absPath (v)
    }
    v, ok = os.LookupEnv ("GOPATH")
    if ok && v != "" {
        // TODO: support possible ':' separator in GOPATH
        v = fp.Join (absPath (v), "src", "github.com", "jrmsdev", "go-jcms", "apps")
    } else {
        v = getEnv ("GOPATH", fp.FromSlash ("/opt/jcms"))
    }
    return absPath (v)
}

func webappDir () string {
    return absPath (fp.Join (baseDir(), WebappName ()))
}

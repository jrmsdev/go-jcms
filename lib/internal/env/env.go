package env

import (
	"os"
	fp "path/filepath"

	"github.com/jrmsdev/go-jcms/lib/internal/logger"
)

var log = logger.New("env")

// OS env defaults - set as JCMS_<UPPERCASE_NAME>
// ie: JCMS_WEBAPP - JCMS_BASEDIR
var webapp = "default"
var basedir = "/opt/jcms"
var datadir = "/var/opt/jcms"

func WebappName() string {
	return getEnv("JCMS_WEBAPP", webapp)
}

func WebappDir() string {
	return absPath(fp.Join(baseDir(), WebappName()))
}

func SettingsFile() string {
	return absPath(fp.Join(WebappDir(), "settings.json"))
}

func DataDir() string {
	d, ok := os.LookupEnv("JCMS_DATADIR")
	d = absPath(d)
	if ok && d != "." {
		return absPath(fp.Join(d, WebappName()))
	}
	return absPath(fp.FromSlash(datadir + "/" + WebappName()))
}

func getEnv(n, d string) string {
	v, isSet := os.LookupEnv(n)
	if isSet {
		return v
	}
	return d
}

func absPath(p string) string {
	rp, err := fp.Abs(fp.Clean(p))
	if err != nil {
		log.Panic("config absPath: %s - %s", p, err)
	}
	return rp
}

func baseDir() string {
	v, ok := os.LookupEnv("JCMS_BASEDIR")
	v = absPath(v)
	if ok && v != "." {
		return v
	}
	return absPath(fp.FromSlash(basedir))
}

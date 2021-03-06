package jcms

import (
	"os"

	"github.com/jrmsdev/go-jcms/lib/internal/asset"
	"github.com/jrmsdev/go-jcms/lib/internal/httpd"
	"github.com/jrmsdev/go-jcms/lib/internal/logger"
	"github.com/jrmsdev/go-jcms/lib/internal/webapps"
	"github.com/jrmsdev/go-jcms/lib/jcms/api"
	"github.com/jrmsdev/go-jcms/lib/jcms/version"
)

var log = logger.New("jcms")
var listening = false
var webappsStarted = false

func init() {
	if os.Getuid() == 0 || os.Geteuid() == 0 {
		panic("do not run as root")
	}
}

func Listen(addr string) string {
	log.Print("version %s", version.String())
	if addr == "" {
		addr = "127.0.0.1:0"
	}
	uri := httpd.Listen(addr)
	log.Print("%s", uri.String())
	listening = true
	return uri.String()
}

func Serve() {
	if listening {
		if !webappsStarted {
			webapps.Start()
			webappsStarted = true
		}
		httpd.Serve()
	} else {
		log.Panic("call jcms.Listen() first")
	}
}

func Stop() {
	if listening {
		httpd.Stop()
		listening = false
	} else {
		log.E("trying to stop a not listening server...")
		log.Panic("jcms.Listen() and jcms.Serve() should be called first")
	}
}

func Logger(tag string) *logger.Logger {
	return logger.New(tag)
}

func LogStart(level string, fh *os.File) error {
	if err := logger.SetLevel(level); err != nil {
		return err
	}
	return logger.File(fh)
}

func LogStop() {
	err := logger.Close()
	if err != nil {
		panic(err)
	}
}

func SetAssetManager(newmanager api.AssetManager) {
	asset.SetManager(newmanager)
}

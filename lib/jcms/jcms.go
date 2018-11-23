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

func Listen() string {
	log.Print("version %s", version.String())
	uri := httpd.Listen()
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
		log.Panic("jcms.Listen() and jcms.Server() should be called first")
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

package webapps

import (
	"net/http"

	"github.com/jrmsdev/go-jcms/lib/internal/app"
	"github.com/jrmsdev/go-jcms/lib/internal/env"
	"github.com/jrmsdev/go-jcms/lib/internal/httpd"
	"github.com/jrmsdev/go-jcms/lib/internal/logger"
	"github.com/jrmsdev/go-jcms/lib/internal/middleware"
	mwloader "github.com/jrmsdev/go-jcms/lib/internal/middleware/base/loader"
	"github.com/jrmsdev/go-jcms/lib/internal/settings"
	apphandler "github.com/jrmsdev/go-jcms/lib/internal/webapps/handler/app"
	"github.com/jrmsdev/go-jcms/lib/internal/webapps/handler/static"
)

var log = logger.New("webapps")

func Start() {
	name := env.WebappName()
	log.D("start %s", name)
	// read settings
	s, err := settings.New(env.SettingsFile())
	if err != nil {
		errHandler(err)
		return
	}
	// TODO: s.Validate()?
	// app
	var a *app.App
	a, err = app.New(name, s)
	if err != nil {
		errHandler(err)
		return
	}
	// middleware loader
	mwloader.Init()
	// middleware enable
	if err = middleware.Enable(s.MiddlewareList); err != nil {
		errHandler(err)
		return
	}
	// app handlers
	static.Handle(a)
	apphandler.Handle(a)
}

func errHandler(err error) {
	status := http.StatusInternalServerError
	log.D("response status: %d", status)
	httpd.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		log.E("INTERNAL ERROR: %s", err)
		http.Error(w, "INTERNAL ERROR: "+err.Error(), status)
	})
}

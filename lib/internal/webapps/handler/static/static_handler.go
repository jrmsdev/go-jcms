package static

import (
	"net/http"
	"path/filepath"

	"github.com/jrmsdev/go-jcms/lib/internal/app"
	"github.com/jrmsdev/go-jcms/lib/internal/env"
	"github.com/jrmsdev/go-jcms/lib/internal/httpd"
	"github.com/jrmsdev/go-jcms/lib/internal/logger"
)

var log = logger.New("static.handler")

func Handle(a *app.App) {
	log.D("app %s", a)
	staticdir := filepath.Join(env.WebappDir(), "static")
	httpd.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir(staticdir))))
}

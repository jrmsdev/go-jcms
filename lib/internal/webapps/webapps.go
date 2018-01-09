package webapps

import (
	"net/http"
	"path/filepath"

	"github.com/jrmsdev/go-jcms/lib/internal/app"
	"github.com/jrmsdev/go-jcms/lib/internal/context/appctx"
	"github.com/jrmsdev/go-jcms/lib/internal/env"
	"github.com/jrmsdev/go-jcms/lib/internal/httpd"
	"github.com/jrmsdev/go-jcms/lib/internal/logger"
	"github.com/jrmsdev/go-jcms/lib/internal/response"
)

var log = logger.New("webapps")

func Start() {
	log.D("start")
	a, err := app.New()
	if err != nil {
		errHandler(err)
		return
	}
	staticHandler(a)
	mainHandler(a)
}

func staticHandler(a *app.App) {
	log.D("static handler: %s", a)
	staticdir := filepath.Join(env.WebappDir(), "static")
	httpd.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir(staticdir))))
}

func mainHandler(a *app.App) {
	log.D("main handler: %s", a)
	httpd.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := appctx.New()
		defer cancel()
		req := r.WithContext(ctx)
		resp := response.New()
		// app handle
		ctx = a.Handle(ctx, resp, req)
		if appctx.Failed(ctx) {
			respError(w, resp)
		} else if appctx.Redirect(ctx) {
			respRedirect(w, resp, req)
		} else {
			writeResp(w, resp)
		}
	})
}

func errHandler(err error) {
	httpd.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.E("INTERNAL ERROR: %s", err)
		http.Error(w, "INTERNAL ERROR: no settings file",
			http.StatusInternalServerError)
	})
}

func respError(w http.ResponseWriter, resp *response.Response) {
	http.Error(w, "ERROR: "+resp.Error(), resp.Status())
}

func respRedirect(
	w http.ResponseWriter,
	resp *response.Response,
	req *http.Request,
) {
	http.Redirect(w, req, resp.Location(), resp.Status())
}

func respHeaders(w http.ResponseWriter, resp *response.Response) {
	log.D("set response headers")
	for h, v := range resp.Headers() {
		w.Header().Set(h, v)
	}
	w.WriteHeader(resp.Status())
}

func writeResp(w http.ResponseWriter, resp *response.Response) {
	respHeaders(w, resp)
	log.D("write response")
	sent, err := w.Write(resp.Body())
	if err != nil {
		log.Panic("write response: %s", err)
	}
	log.D("response sent: %d", sent)
}

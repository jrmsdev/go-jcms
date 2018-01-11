package app

import (
	"net/http"

	"github.com/jrmsdev/go-jcms/lib/internal/app"
	"github.com/jrmsdev/go-jcms/lib/internal/context/appctx"
	"github.com/jrmsdev/go-jcms/lib/internal/httpd"
	"github.com/jrmsdev/go-jcms/lib/internal/logger"
	"github.com/jrmsdev/go-jcms/lib/internal/request"
	"github.com/jrmsdev/go-jcms/lib/internal/response"
)

var log = logger.New("app.handler")

func Handle(a *app.App) {
	log.D("main handler %s", a)
	httpd.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := appctx.New()
		defer cancel()
		req := request.New(ctx, r)
		resp := response.New()
		// app handle
		ctx = a.Handle(ctx, resp, req)
		if appctx.Failed(ctx) {
			respError(w, resp)
		} else if appctx.Redirect(ctx) {
			respRedirect(w, resp, r)
		} else {
			writeResp(w, resp)
		}
	})
}

func respError(w http.ResponseWriter, resp *response.Response) {
	http.Error(w, "ERROR: "+resp.Error(), resp.Status())
}

func respRedirect(
	w http.ResponseWriter,
	resp *response.Response,
	r *http.Request,
) {
	http.Redirect(w, r, resp.Location(), resp.Status())
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
		log.E("write response %s", err)
	}
	log.D("response sent %d", sent)
}

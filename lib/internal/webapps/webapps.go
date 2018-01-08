package webapps

import (
	"log"
	"net/http"

	"github.com/jrmsdev/go-jcms/lib/internal/app"
	"github.com/jrmsdev/go-jcms/lib/internal/context/appctx"
	"github.com/jrmsdev/go-jcms/lib/internal/httpd"
	"github.com/jrmsdev/go-jcms/lib/internal/response"
)

func Start() {
	log.Println("webapps: start")
	a, err := app.New()
	if err != nil {
		errHandler(err)
		return
	}
	mainHandler(a)
}

func mainHandler(a *app.App) {
	log.Println("main handler:", a)
	httpd.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// TODO: use http.Request.WithContext
		resp := response.New()
		ctx, cancel := appctx.New(req, resp)
		defer cancel()
		ctx = a.Handle(ctx)
		if appctx.Failed(ctx) {
			respError(w, resp)
		} else if appctx.Redirect(ctx) {
			respRedirect(w, req, resp)
		} else {
			writeResp(w, resp)
		}
	})
}

func errHandler(err error) {
	httpd.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("INTERNAL ERROR:", err.Error())
		http.Error(w, "INTERNAL ERROR: "+err.Error(),
			http.StatusInternalServerError)
	})
}

func respError(w http.ResponseWriter, resp *response.Response) {
	log.Println("ERROR:", resp.Error())
	http.Error(w, "ERROR: "+resp.Error(), resp.Status())
}

func respRedirect(w http.ResponseWriter, r *http.Request, resp *response.Response) {
	// TODO: redirect response
	//~ http.Redirect (w, r, resp.Location (), resp.Status ())
}

func respHeaders(w http.ResponseWriter, resp *response.Response) {
	for h, v := range resp.Headers() {
		w.Header().Set(h, v)
	}
	w.WriteHeader(resp.Status())
}

func writeResp(w http.ResponseWriter, resp *response.Response) {
	log.Println("write response")
	respHeaders(w, resp)
	sent, err := w.Write(resp.Body())
	if err != nil {
		log.Fatalln("PANIC: write response:", err.Error())
	}
	log.Println("response sent:", sent)
}

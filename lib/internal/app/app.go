package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jrmsdev/go-jcms/lib/internal/context/appctx"
	"github.com/jrmsdev/go-jcms/lib/internal/doctype"
	"github.com/jrmsdev/go-jcms/lib/internal/logger"
	"github.com/jrmsdev/go-jcms/lib/internal/middleware"
	"github.com/jrmsdev/go-jcms/lib/internal/response"
	"github.com/jrmsdev/go-jcms/lib/internal/settings"
	"github.com/jrmsdev/go-jcms/lib/internal/views"

	// init doctype engines
	_ "github.com/jrmsdev/go-jcms/lib/internal/doctype/base/loader"

	// init middleware packages
	_ "github.com/jrmsdev/go-jcms/lib/internal/middleware/base/loader"
)

var log = logger.New("app")

type App struct {
	name     string
	vreg     *views.Registry
	settings *settings.Settings
}

func New(name string, s *settings.Settings) (*App, error) {
	// app object
	a := &App{name, views.Register(s.ViewList), s}
	return a, nil
}

func (a *App) String() string {
	return fmt.Sprintf("app.%s", a.name)
}

func (a *App) Handle(
	ctx context.Context,
	resp *response.Response,
	req *http.Request,
) context.Context {
	// view handler
	view, err := a.vreg.Get(req.URL.Path)
	if err != nil {
		return resp.SetError(ctx, http.StatusNotFound, err.Error())
	}
	// settings reader
	cfg := settings.NewReader(a.settings, view)
	// view redirect
	if view.Redirect != "" {
		return respRedirect(ctx, resp, cfg)
	}
	// middleware PRE
	ctx = middleware.Action(ctx, resp, req, cfg, middleware.ACTION_PRE)
	if appctx.Failed(ctx) {
		return ctx
	}
	// doctype engine
	ctx = doctypeEngine(ctx, resp, req, cfg)
	if appctx.Failed(ctx) {
		return ctx
	}
	// middleware POST
	return middleware.Action(ctx, resp, req, cfg, middleware.ACTION_POST)
}

func respRedirect(
	ctx context.Context,
	resp *response.Response,
	cfg *settings.Reader,
) context.Context {
	var statusmap = map[string]int{
		"permanent": http.StatusPermanentRedirect,
		"temporary": http.StatusTemporaryRedirect,
	}
	status, ok := statusmap[cfg.View.Redirect]
	if !ok {
		return resp.SetError(ctx, http.StatusInternalServerError,
			fmt.Sprintf("invalid view (%s) redirect: %s",
				cfg.View.String(), cfg.View.Redirect))
	}
	if cfg.View.Location == "" {
		cfg.View.Location = "/NOLOCATION"
	}
	return resp.Redirect(ctx, status, cfg.View.Location)
}

func doctypeEngine(
	ctx context.Context,
	resp *response.Response,
	req *http.Request,
	cfg *settings.Reader,
) context.Context {
	log.D("view doctype %s", cfg.View.Doctype)
	eng, err := doctype.GetEngine(cfg.View.Doctype)
	if err != nil {
		return resp.SetError(ctx,
			http.StatusInternalServerError, err.Error())
	}
	log.D("view engine %s", eng.String())
	return eng.Handle(ctx, resp, req, cfg)
}

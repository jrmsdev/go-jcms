package app

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/jrmsdev/go-jcms/lib/internal/context/appctx"
	dbm "github.com/jrmsdev/go-jcms/lib/internal/db"
	"github.com/jrmsdev/go-jcms/lib/internal/doctype"
	"github.com/jrmsdev/go-jcms/lib/internal/env"
	"github.com/jrmsdev/go-jcms/lib/internal/fsutils"
	"github.com/jrmsdev/go-jcms/lib/internal/logger"
	"github.com/jrmsdev/go-jcms/lib/internal/middleware"
	"github.com/jrmsdev/go-jcms/lib/internal/request"
	"github.com/jrmsdev/go-jcms/lib/internal/response"
	"github.com/jrmsdev/go-jcms/lib/internal/settings"
	"github.com/jrmsdev/go-jcms/lib/internal/views"

	// init db engines
	_ "github.com/jrmsdev/go-jcms/lib/internal/db/loader"

	// init doctype engines
	_ "github.com/jrmsdev/go-jcms/lib/internal/doctype/base/loader"
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
	req *request.Request,
) context.Context {
	// TODO: connect to database
	_, dberr := dbm.Open("memdb://jcmsdb")
	if dberr != nil {
		log.E("db open %s", dberr.Error())
		return resp.SetError(ctx, http.StatusInternalServerError,
			"ERROR: database connection")
	}
	// view handler
	view, err := a.vreg.Get(req.URL.Path)
	if err != nil {
		return resp.SetError(ctx, http.StatusNotFound, err.Error())
	}
	// settings reader
	cfg := settings.NewReader(a.settings, view)
	// view redirect
	if view.Redirect != "" {
		return viewRedirect(ctx, resp, cfg)
	}
	// middleware PRE
	ctx = middleware.Action(ctx, resp, req, cfg, middleware.ACTION_PRE)
	if appctx.Failed(ctx) {
		return ctx
	}
	// doctype engine
	ctx = doctypeEngine(ctx, resp, req, cfg)
	if appctx.Failed(ctx) || appctx.EngineFailed(ctx) {
		return ctx
	}
	// middleware POST
	return middleware.Action(ctx, resp, req, cfg, middleware.ACTION_POST)
}

func viewRedirect(
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
	req *request.Request,
	cfg *settings.Reader,
) context.Context {
	log.D("view doctype %s", cfg.View.Doctype)
	eng, err := doctype.GetEngine(cfg.View.Doctype)
	if err != nil {
		return resp.SetError(ctx,
			http.StatusInternalServerError, err.Error())
	}
	docroot := filepath.Join(env.WebappDir(), "docroot")
	if !fsutils.DirExists(docroot) {
		log.E("docroot not found:", docroot)
		return resp.SetError(ctx,
			http.StatusInternalServerError, "docroot not found")
	}
	log.D("%s handle", eng)
	ctx = eng.Handle(ctx, resp, req, cfg, docroot)
	if appctx.Failed(ctx) {
		ctx = appctx.EngineFail(ctx)
		return eng.HandleError(ctx, resp, req, cfg, docroot)
	}
	return ctx
}

package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jrmsdev/go-jcms/lib/internal/context/appctx"
	"github.com/jrmsdev/go-jcms/lib/internal/doctype"
	"github.com/jrmsdev/go-jcms/lib/internal/env"
	"github.com/jrmsdev/go-jcms/lib/internal/fsutils"
	"github.com/jrmsdev/go-jcms/lib/internal/logger"
	"github.com/jrmsdev/go-jcms/lib/internal/middleware"
	"github.com/jrmsdev/go-jcms/lib/internal/response"
	"github.com/jrmsdev/go-jcms/lib/internal/views"

	// init doctype engines
	_ "github.com/jrmsdev/go-jcms/lib/internal/doctype/base/loader"

	// init middleware packages
	_ "github.com/jrmsdev/go-jcms/lib/internal/middleware/base/loader"
)

var log = logger.New("app")

type App struct {
	name string
	vreg *views.Registry
}

func New() (*App, error) {
	name := env.WebappName()
	log.D(name)
	s, err := getSettings()
	if err != nil {
		return nil, err
	}
	a := &App{name, views.Register(s.Views)}
	// middleware enable
	if err := middleware.Enable(s.Middleware); err != nil {
		return nil, err
	}
	return a, nil
}

func getSettings() (*Settings, error) {
	fn := env.SettingsFile()
	log.D(fn)
	if !fsutils.FileExists(fn) {
		return nil, fmt.Errorf("file not found: %s", fn)
	}
	return readSettings(fn)
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
	// view redirect
	if view.Redirect != "" {
		return respRedirect(ctx, resp, view)
	}
	// middleware PRE
	ctx = middleware.Action(ctx, resp, middleware.ACTION_PRE, req)
	if appctx.Failed(ctx) {
		return ctx
	}
	// doctype engine
	ctx = doctypeEngine(ctx, resp, view, req)
	if appctx.Failed(ctx) {
		return ctx
	}
	// middleware POST
	return middleware.Action(ctx, resp, middleware.ACTION_POST, req)
}

func respRedirect(
	ctx context.Context,
	resp *response.Response,
	view *views.View,
) context.Context {
	var statusmap = map[string]int{
		"permanent": http.StatusPermanentRedirect,
		"temporary": http.StatusTemporaryRedirect,
	}
	status, ok := statusmap[view.Redirect]
	if !ok {
		return resp.SetError(ctx, http.StatusInternalServerError,
			fmt.Sprintf("invalid view (%s) redirect: %s",
				view.String(), view.Redirect))
	}
	if view.Location == "" {
		view.Location = "/NOLOCATION"
	}
	return resp.Redirect(ctx, status, view.Location)
}

func doctypeEngine(
	ctx context.Context,
	resp *response.Response,
	view *views.View,
	req *http.Request,
) context.Context {
	log.D("view doctype", view.Doctype)
	eng, err := doctype.GetEngine(view.Doctype)
	if err != nil {
		return resp.SetError(ctx,
			http.StatusInternalServerError, err.Error())
	}
	log.D("view engine: %s", eng.String())
	return eng.Handle(ctx, resp, view, req)
}

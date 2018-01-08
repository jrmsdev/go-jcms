package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/jrmsdev/go-jcms/lib/internal/context/appctx"
	"github.com/jrmsdev/go-jcms/lib/internal/doctype"
	"github.com/jrmsdev/go-jcms/lib/internal/env"
	"github.com/jrmsdev/go-jcms/lib/internal/fsutils"
	"github.com/jrmsdev/go-jcms/lib/internal/response"
	"github.com/jrmsdev/go-jcms/lib/internal/views"
)

type App struct {
	name string
	vreg *views.Registry
}

func New() (*App, error) {
	name := env.WebappName()
	log.Println("app:", name)
	s, err := getSettings()
	if err != nil {
		return nil, err
	}
	a := &App{name, views.Register(s.Views)}
	return a, nil
}

func getSettings() (*Settings, error) {
	fn := env.SettingsFile()
	log.Println("app:", fn)
	if !fsutils.FileExists(fn) {
		return nil, fmt.Errorf("file not found: %s", fn)
	}
	return readSettings(fn)
}

func (a *App) String() string {
	return fmt.Sprintf("app.%s", a.name)
}

func (a *App) Handle(
	req *http.Request,
	resp *response.Response,
) context.Context {
	ctx := req.Context()
	view, err := a.vreg.Get(req.URL.Path)
	if err != nil {
		resp.SetError(http.StatusNotFound, err.Error())
		return appctx.Fail(ctx)
	}
	if view.Redirect != "" {
		err := viewRedirect(view, resp)
		if err != nil {
			resp.SetError(http.StatusInternalServerError,
				err.Error())
			return appctx.Fail(ctx)
		}
		return appctx.SetRedirect(ctx)
	}
	return doctypeEngine(ctx, view, req, resp)
}

func viewRedirect(view *views.View, resp *response.Response) error {
	var statusmap = map[string]int{
		"permanent": http.StatusPermanentRedirect,
		"temporary": http.StatusTemporaryRedirect,
	}
	status, ok := statusmap[view.Redirect]
	if !ok {
		return fmt.Errorf("invalid view (%s) redirect: %s",
			view.String(), view.Redirect)
	}
	if view.Location == "" {
		view.Location = "/NOLOCATION"
	}
	resp.Redirect(status, view.Location)
	return nil
}

func doctypeEngine(
	ctx context.Context,
	view *views.View,
	req *http.Request,
	resp *response.Response,
) context.Context {
	log.Println("app: view doctype", view.Doctype)
	eng, err := doctype.GetEngine(view.Doctype)
	if err != nil {
		resp.SetError(http.StatusInternalServerError, err.Error())
		ctx = appctx.Fail(ctx)
		return ctx
	}
	log.Println("app: view engine", eng.String())
	return eng.Handle(req, resp)
}

package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/jrmsdev/go-jcms/lib/internal/context/appctx"
	"github.com/jrmsdev/go-jcms/lib/internal/doctype"
	"github.com/jrmsdev/go-jcms/lib/internal/env"
	"github.com/jrmsdev/go-jcms/lib/internal/utils"
	"github.com/jrmsdev/go-jcms/lib/internal/views"
)

type App struct {
	name     string
	settings *Settings
	views    *views.Registry
}

func New() (*App, error) {
	name := env.WebappName()
	log.Println("app:", name)
	s, err := getSettings()
	if err != nil {
		return nil, err
	}
	reg := views.Register(s.Views)
	a := &App{name, s, reg}
	return a, nil
}

func (a *App) String() string {
	return fmt.Sprintf("app.%s", a.name)
}

func (a *App) Handle(ctx context.Context) context.Context {
	var (
		err  error
		view *views.View
		eng  doctype.Engine
	)
	resp := appctx.Response(ctx)
	req := appctx.Request(ctx)

	view, err = a.findView(req.URL.Path)
	if err != nil {
		resp.SetError(http.StatusNotFound, err.Error())
		ctx = appctx.Fail(ctx)
		return ctx
	}

	log.Println("app: view doctype", view.Doctype)
	eng, err = doctype.GetEngine(view.Doctype)
	if err != nil {
		resp.SetError(http.StatusInternalServerError, err.Error())
		ctx = appctx.Fail(ctx)
		return ctx
	}

	log.Println("app: view engine", eng.String())
	return eng.Handle(ctx)
}

func (a *App) findView(path string) (*views.View, error) {
	return a.views.Get(path)
}

func getSettings() (*Settings, error) {
	fn := env.SettingsFile()
	log.Println("app:", fn)
	if !utils.FileExists(fn) {
		return nil, fmt.Errorf("file not found: %s", fn)
	}
	return readSettings(fn)
}

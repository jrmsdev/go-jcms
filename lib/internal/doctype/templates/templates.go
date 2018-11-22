package templates

import (
	"context"
	"html/template"
	"net/http"
	"path"
	"path/filepath"

	"github.com/jrmsdev/go-jcms/lib/internal/asset"
	"github.com/jrmsdev/go-jcms/lib/internal/doctype"
	"github.com/jrmsdev/go-jcms/lib/internal/doctype/base"
	"github.com/jrmsdev/go-jcms/lib/internal/doctype/templates/funcs"
	"github.com/jrmsdev/go-jcms/lib/internal/logger"
	"github.com/jrmsdev/go-jcms/lib/internal/request"
	"github.com/jrmsdev/go-jcms/lib/internal/response"
	"github.com/jrmsdev/go-jcms/lib/internal/settings"
)

var log = logger.New("doctype.templates")

func init() {
	doctype.Register("templates", newEngine())
}

type engine struct {
	base.Engine
}

func newEngine() *engine {
	return &engine{base.New("templates")}
}

func (e *engine) Handle(
	ctx context.Context,
	resp *response.Response,
	req *request.Request,
	cfg *settings.Reader,
	docroot string,
) context.Context {
	var (
		maintplfn string
		viewtplfn string
	)
	log.D("docroot %s", docroot)
	args := cfg.View.Args
	// get template files
	layout := args.Get("layout", "main").String()
	maintplfn = getMainTpl(cfg, docroot, layout)
	viewtplfn = getViewTpl(cfg, docroot, req.URL.Path)
	// templates data
	tpldata := newData()
	return tplHandle(ctx, resp, req, cfg, docroot,
		maintplfn, viewtplfn, tpldata)
}

func (e *engine) HandleError(
	ctx context.Context,
	resp *response.Response,
	req *request.Request,
	cfg *settings.Reader,
	docroot string,
) context.Context {
	var (
		maintplfn string
		viewtplfn string
	)
	// get error templates
	maintplfn = getMainTpl(cfg, docroot, "error")
	viewtplfn = getViewTpl(cfg, docroot, "error")
	// templates data
	tpldata := newErrorData()
	return tplHandle(ctx, resp, req, cfg, docroot,
		maintplfn, viewtplfn, tpldata)
}

func tplHandle(
	ctx context.Context,
	resp *response.Response,
	_ *request.Request,
	_ *settings.Reader,
	docroot string,
	maintplfn string,
	viewtplfn string,
	tpldata *Data,
) context.Context {
	var (
		err     error
		tplname string
		maintpl *template.Template
		viewtpl *template.Template
	)
	// parse main template
	maintpl, err = parseMainTpl(docroot, maintplfn)
	if err != nil {
		log.E("parse main template: %s", err.Error())
		return resp.SetError(ctx, http.StatusInternalServerError,
			"ERROR: parse main template")
	}
	// parse view template (if provided)
	if viewtplfn != "" {
		viewtpl, err = parseViewTpl(maintpl, viewtplfn)
		if err != nil {
			log.E("parse view template: %s", err.Error())
			return resp.SetError(ctx, http.StatusInternalServerError,
				"ERROR: parse view template")
		}
	}
	// execute template
	resp.SetTemplateLayout(tplName(docroot, maintplfn))
	tplname = tplName(docroot, viewtplfn)
	resp.SetTemplate(tplname)
	log.D("exec %s", tplname)
	err = viewtpl.Execute(resp, tpldata)
	if err != nil {
		log.E("exec template: %s", err.Error())
		return resp.SetError(ctx, http.StatusInternalServerError,
			"ERROR: exec template")
	}
	resp.SetStatus(http.StatusOK)
	return ctx
}

func getMainTpl(cfg *settings.Reader, docroot, layout string) string {
	return filepath.Join(docroot, layout+".tpl")
}

func getViewTpl(cfg *settings.Reader, docroot string, fn string) string {
	fn = path.Clean(fn)
	if fn == "" || fn == "/" {
		fn = "index"
	}
	return filepath.Join(docroot, fn+".html")
}

func parseMainTpl(docroot, fn string) (*template.Template, error) {
	content, err := asset.ReadFile(fn)
	if err != nil {
		return nil, err
	}
	name := tplName(docroot, fn)
	return template.New(name).Funcs(funcs.Map).Parse(string(content))
}

func parseViewTpl(main *template.Template, fn string) (*template.Template, error) {
	content, err := asset.ReadFile(fn)
	if err != nil {
		return nil, err
	}
	return template.Must(main.Clone()).Parse(string(content))
}

func tplName(docroot, filename string) string {
	n, err := filepath.Rel(docroot, filename)
	if err != nil {
		n = "ERROR:" + err.Error()
	}
	return n
}

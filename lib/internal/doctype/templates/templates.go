package templates

import (
	"context"
	"html/template"
	"io/ioutil"
	"net/http"
	"path"
	"path/filepath"

	"github.com/jrmsdev/go-jcms/lib/internal/doctype"
	"github.com/jrmsdev/go-jcms/lib/internal/doctype/base"
	"github.com/jrmsdev/go-jcms/lib/internal/doctype/templates/funcs"
	"github.com/jrmsdev/go-jcms/lib/internal/fsutils"
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
		ok        bool
		err       error
		tplname   string
		maintplfn string
		viewtplfn string
		maintpl   *template.Template
		viewtpl   *template.Template
	)
	log.D("docroot", docroot)
	// get template files
	maintplfn, ok = getMainTpl(cfg, docroot)
	if !ok {
		log.E("main template not found:", maintplfn)
		return resp.SetError(ctx, http.StatusInternalServerError,
			"main template not found")
	}
	viewtplfn, ok = getViewTpl(cfg, req, docroot)
	if !ok {
		log.E("view template not found:", viewtplfn)
		return resp.SetError(ctx, http.StatusNotFound, "not found")
	}
	// parse templates
	maintpl, err = parseMainTpl(maintplfn)
	if err != nil {
		log.E("parse main template:", err.Error())
		return resp.SetError(ctx, http.StatusInternalServerError,
			"ERROR: parse main template")
	}
	viewtpl, err = parseViewTpl(maintpl, viewtplfn)
	if err != nil {
		log.E("parse view template:", err.Error())
		return resp.SetError(ctx, http.StatusInternalServerError,
			"ERROR: parse view template")
	}
	// templates data
	tpldata := newData()
	// execute main template
	tplname = tplName(docroot, maintplfn)
	resp.SetTemplateLayout(tplname)
	log.D("exec main", tplname)
	err = execTpl(resp, maintpl, tpldata)
	if err != nil {
		log.E("exec main template:", err.Error())
		return resp.SetError(ctx, http.StatusInternalServerError,
			"ERROR: exec main template")
	}
	// execute view template
	tplname = tplName(docroot, viewtplfn)
	resp.SetTemplate(tplname)
	log.D("exec view", tplname)
	err = execTpl(resp, viewtpl, tpldata)
	if err != nil {
		log.E("exec view template:", err.Error())
		return resp.SetError(ctx, http.StatusInternalServerError,
			"ERROR: exec view template")
	}
	resp.SetStatus(http.StatusOK)
	return ctx
}

func getMainTpl(cfg *settings.Reader, docroot string) (string, bool) {
	filename := filepath.Join(docroot, "main.tpl")
	if !fsutils.FileExists(filename) {
		return filename, false
	}
	return filename, true
}

func getViewTpl(
	cfg *settings.Reader,
	req *request.Request,
	docroot string,
) (string, bool) {
	fn := req.URL.Path
	if fn == "" || fn == "/" {
		fn = path.Clean(cfg.View.Path)
	}
	if fn == "" || fn == "/" {
		fn = "index"
	}
	filename := filepath.Join(docroot, fn+".html")
	if !fsutils.FileExists(filename) {
		return filename, false
	}
	return filename, true
}

func parseMainTpl(fn string) (*template.Template, error) {
	content, err := ioutil.ReadFile(fn)
	if err != nil {
		return nil, err
	}
	return template.New("maintpl").Funcs(funcs.Map).Parse(string(content))
}

func parseViewTpl(main *template.Template, fn string) (*template.Template, error) {
	content, err := ioutil.ReadFile(fn)
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

func execTpl(
	resp *response.Response,
	tpl *template.Template,
	data *Data,
) error {
	return tpl.Execute(resp, data)
}

package templates

import (
	//~ "net/http"
	"testing"

	//~ "github.com/jrmsdev/go-jcms/lib/internal/context/appctx"
	"github.com/jrmsdev/go-jcms/lib/internal/doctype"
	"github.com/jrmsdev/go-jcms/lib/internal/testing/testeng"
)

const testengName = "templates"

func TestEngine(t *testing.T) {
	e, err := doctype.GetEngine(testengName)
	if err != nil {
		t.Fatal(err)
	}
	if e.Type() != testengName {
		t.Error(".Type !=", testengName)
	}
}

func getResult(t *testing.T, path string) *testeng.Result {
	return testeng.Handle(t, testengName,
		&testeng.Query{
			App:  testengName,
			Path: path,
		})
}

//~ func TestHandle(t *testing.T) {
	//~ r := getResult(t, "/test")
	//~ if appctx.Failed(r.Ctx) {
		//~ t.Error("handle context should not fail:", r.Resp.Error())
	//~ }
	//~ status := r.Resp.Status()
	//~ if status != http.StatusOK {
		//~ t.Error("invalid resp status:", status)
		//~ t.Log("expect:", http.StatusOK)
	//~ }
	//~ if r.Resp.Template() != "test.html" {
		//~ t.Error("invalid resp template:", r.Resp.Template())
		//~ t.Log("expect: test.html")
	//~ }
	//~ if r.Resp.TemplateLayout() != "main.tpl" {
		//~ t.Error("invalid resp template layout:", r.Resp.TemplateLayout())
		//~ t.Log("expect: main.tpl")
	//~ }
//~ }

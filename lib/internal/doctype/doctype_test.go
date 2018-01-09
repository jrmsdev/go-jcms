package doctype

import (
	"context"
	"net/http"
	"testing"

	"github.com/jrmsdev/go-jcms/lib/internal/doctype/base"
	"github.com/jrmsdev/go-jcms/lib/internal/response"
	"github.com/jrmsdev/go-jcms/lib/internal/views"
)

type testEngine struct {
	base.Engine
}

func newTestEngine() *testEngine {
	return &testEngine{base.New("testengine")}
}

func (e *testEngine) Handle(
	_ *views.View,
	req *http.Request,
	_ *response.Response,
) context.Context {
	return req.Context()
}

func TestRegister(t *testing.T) {
	Register("testengine", newTestEngine())
	e, err := GetEngine("testengine")
	if err != nil {
		t.Fatal("get engine error:", err)
	}
	et := e.Type()
	if et != "testengine" {
		t.Error("invalid engine type:", et)
	}
	es := e.String()
	if es != "doctype.testengine" {
		t.Error("invalid engine string:", es)
	}
	_, err = GetEngine("notexistent")
	if err == nil {
		t.Fatal("get not existent engine should have failed")
	}
}

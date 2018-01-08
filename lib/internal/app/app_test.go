package app

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jrmsdev/go-jcms/lib/internal/context/appctx"
)

func TestNewApp(t *testing.T) {
	testappEnv("testing")
	defer testappEnv("") // cleanup
	a, err := New()
	if err != nil {
		t.Fatal(err)
	}
	if a.String() != fmt.Sprintf("app.%s", a.name) {
		t.Error("a.String != app.<name>")
	}
}

func TestNewAppSettingsError(t *testing.T) {
	testappEnv("invalidapp")
	defer testappEnv("") // cleanup
	a, err := New()
	if err == nil {
		t.Log(a, err)
		t.Error("settings file for invalidapp should fail")
	}
}

func TestAppHandle(t *testing.T) {
	tapp := newTestApp()
	r := tapp.Handle("/test")
	if appctx.Failed(r.Ctx) {
		t.Log(r.Resp.Error())
		t.Error("app.Handle should not fail")
	}
	body := strings.TrimSpace(string(r.Resp.Body()))
	if body != "testing" {
		t.Error("invalid resp body:", body)
	}
}

func TestAppHandleViewNotFound(t *testing.T) {
	tapp := newTestApp()
	r := tapp.Handle("/test/view.not.found")
	if !appctx.Failed(r.Ctx) {
		t.Fatal("app.Handle should fail")
	}
	if r.Resp.Error() != "view: not found: /test/view.not.found" {
		t.Log(r.Resp.Error())
		t.Error("wrong app.Handle view not found error message")
	}
}

func TestAppHandleInvalidEngine(t *testing.T) {
	tapp := newTestApp()
	r := tapp.Handle("/test/doctype.engine.invalid")
	if !appctx.Failed(r.Ctx) {
		t.Fatal("app.Handle should fail")
	}
	if r.Resp.Error() != "invalid doctype engine: invalid.engine" {
		t.Log(r.Resp.Error())
		t.Error("wrong app.Handle invalid doctype engine error message")
	}
}

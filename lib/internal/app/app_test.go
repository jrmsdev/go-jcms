package app

import (
	"net/http"
	"strings"
	"testing"

	"github.com/jrmsdev/go-jcms/lib/internal/context/appctx"
)

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

func TestAppHandleRedirectPermanent(t *testing.T) {
	tapp := newTestApp()
	r := tapp.Handle("/test/redirect/permanent")
	if !appctx.Redirect(r.Ctx) {
		t.Error("no redirect context")
	}
	location := r.Resp.Location()
	if location != "/test/redirect/location" {
		t.Error("invalid redirect location:", location)
	}
	status := r.Resp.Status()
	if status != http.StatusPermanentRedirect {
		t.Error("invalid redirect status:", status)
	}
}

func TestAppHandleRedirectTemporary(t *testing.T) {
	tapp := newTestApp()
	r := tapp.Handle("/test/redirect/temporary")
	if !appctx.Redirect(r.Ctx) {
		t.Error("no redirect context")
	}
	location := r.Resp.Location()
	if location != "/test/redirect/location" {
		t.Error("invalid redirect location:", location)
	}
	status := r.Resp.Status()
	if status != http.StatusTemporaryRedirect {
		t.Error("invalid redirect status:", status)
	}
}

func TestAppHandleRedirectInvalidStatus(t *testing.T) {
	tapp := newTestApp()
	r := tapp.Handle("/test/redirect/invalidstatus")
	if !appctx.Failed(r.Ctx) {
		t.Error("redirect with invalid status should have failed")
	}
	rstat := r.Resp.Status()
	if rstat != http.StatusInternalServerError {
		t.Error("invalid response status:", rstat)
	}
}

func TestAppHandleRedirectNoLocation(t *testing.T) {
	tapp := newTestApp()
	r := tapp.Handle("/test/redirect/nolocation")
	if !appctx.Redirect(r.Ctx) {
		t.Error("no redirect context")
	}
	location := r.Resp.Location()
	if location != "/NOLOCATION" {
		t.Error("invalid default location:", location)
	}
}

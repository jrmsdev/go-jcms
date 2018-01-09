package response

import (
	"net/http"
	"testing"
)

func TestHeaders(t *testing.T) {
	r := New()
	r.SetHeader("x-test-header", "testing")
	h := r.Headers()
	v, found := h["x-test-header"]
	t.Log(v, found)
	if !found {
		t.Error("response header not found")
	}
	if v != "testing" {
		t.Error("invalid response header:", v)
	}
}

func TestStatus(t *testing.T) {
	r := New()
	status := r.Status()
	if status != http.StatusNotImplemented {
		t.Error("invalid default status:", status)
	}
	r.SetStatus(999)
	status = r.Status()
	if status != 999 {
		t.Error("set status failed:", status)
	}
}

func TestError(t *testing.T) {
	r := New()
	errmsg := r.Error()
	if errmsg != "NOERRMSG" {
		t.Error("invalid default error message:", errmsg)
	}
	r.SetError(999, "error message")
	status := r.Status()
	if status != 999 {
		t.Error("set error status failed:", status)
	}
	errmsg = r.Error()
	if errmsg != "error message" {
		t.Error("set error message failed:", errmsg)
	}
}

func TestRedirect(t *testing.T) {
	r := New()
	location := r.Location()
	if location != "NOLOCATION" {
		t.Error("invalid default location:", location)
	}
	r.Redirect(999, "/redirect/location")
	status := r.Status()
	if status != 999 {
		t.Error("set redirect status failed:", status)
	}
	location = r.Location()
	if location != "/redirect/location" {
		t.Error("set redirect location failed:", location)
	}
}

func TestBody(t *testing.T) {
	r := New()
	body := string(r.Body())
	if body != "" {
		t.Error("body should be empty at init:", body)
	}
	n, err := r.Write([]byte("testing body"))
	if n != 12 {
		t.Error("invalid write length:", n)
	}
	if err != nil {
		t.Error("body write failed:", err)
		return
	}
	body = string(r.Body())
	if body != "testing body" {
		t.Error("write body failed:", body)
	}
}

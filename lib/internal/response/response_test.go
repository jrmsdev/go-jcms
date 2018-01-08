package response

import (
	"testing"
)

func TestSetHeader(t *testing.T) {
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

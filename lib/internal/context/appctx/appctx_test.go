package appctx

import (
	"context"
	"net/http"
	"testing"
)

func TestCtx(t *testing.T) {
	var cancel context.CancelFunc
	req := &http.Request{}
	req, cancel = New(req)
	defer cancel()
	ctx := req.Context()
	testFailed(t, ctx)
	testRedirect(t, ctx)
}

func testFailed(t *testing.T, ctx context.Context) {
	if Failed(ctx) {
		t.Error("ctx initiated in failed status")
	}
	if !Failed(Fail(ctx)) {
		t.Error("ctx is not in failed status")
	}
}

func testRedirect(t *testing.T, ctx context.Context) {
	if Redirect(ctx) {
		t.Error("ctx initiated in redirect status")
	}
	if !Redirect(SetRedirect(ctx)) {
		t.Error("ctx is not in failed status")
	}
}

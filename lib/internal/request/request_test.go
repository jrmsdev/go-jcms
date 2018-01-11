package request

import (
	"context"
	"net/http"
	"net/url"
	"testing"
)

func getReq(path string) *Request {
	r := &http.Request{}
	r.URL, _ = url.Parse("http://127.0.0.1:0" + path)
	return New(context.Background(), r)
}

func TestRequest(t *testing.T) {
	req := getReq("/test")
	if req.URL.Path != "/test" {
		t.Error("invalid req path:", req.URL.Path, " - expect: /test")
	}
}

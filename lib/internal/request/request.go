package request

import (
	"context"
	"net/http"
)

type Request struct {
	*http.Request
}

func New(ctx context.Context, req *http.Request) *Request {
	return &Request{req.WithContext(ctx)}
}

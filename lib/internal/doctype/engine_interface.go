package doctype

import (
	"context"
	"net/http"

	"github.com/jrmsdev/go-jcms/lib/internal/response"
	"github.com/jrmsdev/go-jcms/lib/internal/views"
)

type Engine interface {
	Type() string
	String() string
	Handle(context.Context, *response.Response,
		*views.View, *http.Request) context.Context
}

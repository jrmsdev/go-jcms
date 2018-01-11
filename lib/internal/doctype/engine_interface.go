package doctype

import (
	"context"
	"net/http"

	"github.com/jrmsdev/go-jcms/lib/internal/response"
	"github.com/jrmsdev/go-jcms/lib/internal/settings"
)

type Engine interface {
	Type() string
	String() string
	Handle(context.Context, *response.Response, *http.Request,
		*settings.Reader, string) context.Context
}

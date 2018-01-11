package doctype

import (
	"context"

	"github.com/jrmsdev/go-jcms/lib/internal/request"
	"github.com/jrmsdev/go-jcms/lib/internal/response"
	"github.com/jrmsdev/go-jcms/lib/internal/settings"
)

type Engine interface {
	Type() string
	String() string
	Handle(context.Context, *response.Response, *request.Request,
		*settings.Reader, string) context.Context
}

package doctype

import (
	"context"
	"net/http"

	"github.com/jrmsdev/go-jcms/lib/internal/response"
)

type Engine interface {
	Type() string
	String() string
	Handle(*http.Request, *response.Response) context.Context
}

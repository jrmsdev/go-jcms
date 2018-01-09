package middleware

import (
	"context"
	"net/http"

	"github.com/jrmsdev/go-jcms/lib/internal/context/appctx"
	"github.com/jrmsdev/go-jcms/lib/internal/logger"
	"github.com/jrmsdev/go-jcms/lib/internal/response"
)

var log = logger.New("middleware")

type MiddlewareAction int

const (
	ACTION_PRE MiddlewareAction = iota
	ACTION_POST
)

type Middleware interface {
	Name() string
	Action(
		ctx context.Context,
		resp *response.Response,
		action MiddlewareAction,
		req *http.Request,
	) context.Context
}

var mwmap = make(map[string]Middleware)
var actiondb = make(map[MiddlewareAction][]string)

func init() {
	actiondb[ACTION_PRE] = make([]string, 0)
	actiondb[ACTION_POST] = make([]string, 0)
}

func Register(mw Middleware, actions ...MiddlewareAction) {
	name := mw.Name()
	_, exists := mwmap[name]
	if exists {
		log.Panic("already registered: %s", name)
	}
	for _, act := range actions {
		actiondb[act] = append(actiondb[act], name)
	}
	mwmap[name] = mw
}

func Enable(settings []*Settings) error {
	// TODO: middleware.Enable
	return nil
}

func Action(
	ctx context.Context,
	resp *response.Response,
	action MiddlewareAction,
	req *http.Request,
) context.Context {
	for _, name := range actiondb[action] {
		mw := mwmap[name]
		ctx = mw.Action(ctx, resp, action, req)
		if appctx.Failed(ctx) {
			return ctx
		}
	}
	return ctx
}

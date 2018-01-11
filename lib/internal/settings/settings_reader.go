package settings

import (
	"fmt"

	"github.com/jrmsdev/go-jcms/lib/internal/settings/middleware"
	"github.com/jrmsdev/go-jcms/lib/internal/settings/view"
)

type Reader struct {
	View       *view.Settings
	Middleware *middleware.Settings
	mwmap      map[string]*middleware.Settings
}

func NewReader(
	src *Settings,
	v *view.Settings,
) *Reader {
	mwmap := make(map[string]*middleware.Settings)
	for _, mw := range src.MiddlewareList {
		mwmap[mw.Name] = mw
	}
	return &Reader{View: v, mwmap: mwmap}
}

func (r *Reader) Reset() {
	r.View.Args.SetPrefix("")
	r.Middleware = nil
}

func (r *Reader) SetMiddleware(name string) error {
	var mw Section
	var ok bool
	mw, ok = r.mwmap[name]
	if !ok {
		return fmt.Errorf("set middleware invalid name: %s", name)
	}
	r.View.Args.SetPrefix(mw.ID())
	r.Middleware = mw.(*middleware.Settings)
	return nil
}

package views

import (
	"fmt"
	xpath "path"

	"github.com/jrmsdev/go-jcms/lib/internal/logger"
	"github.com/jrmsdev/go-jcms/lib/internal/settings/view"
)

var log = logger.New("views.registry")

type Registry struct {
	db  map[string]*view.Settings
	idx map[string]string // paths index for faster access (I hope)
}

func Register(vlist []*view.Settings) *Registry {
	r := &Registry{}
	r.db = make(map[string]*view.Settings)
	r.idx = make(map[string]string)
	for _, v := range vlist {
		// clean view path, it comes from settings.xml file
		v.Path = xpath.Clean(v.Path)
		// TODO: check duplicate view names and / or paths
		r.db[v.Name] = v
		r.idx[v.Path] = v.Name
	}
	return r
}

func (r *Registry) Get(path string) (*view.Settings, error) {
	idx, found := r.idx[path]
	if !found {
		log.E("not found: %s", path)
		return nil, fmt.Errorf("not found: %s", path)
	}
	v := r.db[idx]
	if v.UseView != "" {
		return r.useView(path, v.UseView)
	}
	return v, nil
}

func (r *Registry) useView(path, name string) (*view.Settings, error) {
	log.D("useview: %s", name)
	v, found := r.db[name]
	if !found {
		log.E("useview not found: %s", name)
		return nil, fmt.Errorf("not found: %s", path)
	}
	return v, nil
}

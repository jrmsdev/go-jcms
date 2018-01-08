package views

import (
	"fmt"
	"log"
	xpath "path"
)

type Registry struct {
	db  map[string]*View
	idx map[string]string // paths index for faster access (I hope)
}

func Register(vlist []*View) *Registry {
	r := &Registry{}
	r.db = make(map[string]*View)
	r.idx = make(map[string]string)
	for _, v := range vlist {
		// clean view path, it comes from settings.xml file
		v.Path = xpath.Clean(v.Path)
		// TODO: check duplicate view names
		r.db[v.Name] = v
		r.idx[v.Path] = v.Name
	}
	return r
}

func (r *Registry) Get(path string) (*View, error) {
	idx, found := r.idx[path]
	if !found {
		log.Println("view: not found:", path)
		return nil, fmt.Errorf("view: not found: %s", path)
	}
	v := r.db[idx]
	if v.UseView != "" {
		return r.useView(v.UseView)
	}
	return v, nil
}

func (r *Registry) useView(name string) (*View, error) {
	log.Println("view: useview", name)
	v, found := r.db[name]
	if !found {
		log.Println("view: useview not found:", name)
		return nil, fmt.Errorf("view: useview not found: %s", name)
	}
	return v, nil
}

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
		v.Path = xpath.Clean(v.Path) // clean path, it comes from settings.xml
		r.db[v.Name] = v
		r.idx[v.Path] = v.Name
	}
	return r
}

func (r *Registry) Get(path string) (*View, error) {
	idx, found := r.idx[path]
	if !found {
		log.Println("view: not found", path)
		return nil, fmt.Errorf("view: not found %s", path)
	}
	return r.db[idx], nil
}

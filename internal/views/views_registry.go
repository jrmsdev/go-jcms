package views

import (
    "path"
)

type Registry struct {
    db map[string]*View
    idx map[string]string // paths index for faster access (I hope)
}

func Register (vlist []*View) *Registry {
    r := &Registry{}
    r.db = make (map[string]*View)
    r.idx = make (map[string]string)
    for _, v := range vlist {
        v.Path = path.Clean (v.Path) // clean path, it comes from settings.xml
        r.db[v.Name] = v
        r.idx[v.Path] = v.Name
    }
    return r
}

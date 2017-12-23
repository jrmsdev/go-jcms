package views

type Registry struct {
    db []*View
}

func Register (vlist []*View) *Registry {
    return &Registry{vlist}
}

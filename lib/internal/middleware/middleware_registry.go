package middleware

import (
	"fmt"
)

type mwRegistry struct {
	db       map[string]Middleware
	actiondb map[MiddlewareAction][]string
	enable   map[MiddlewareAction][]string
}

func newRegistry() *mwRegistry {
	return &mwRegistry{
		db:       make(map[string]Middleware),
		actiondb: make(map[MiddlewareAction][]string),
		enable:   make(map[MiddlewareAction][]string),
	}
}

func (r *mwRegistry) Register(mw Middleware, actions ...MiddlewareAction) {
	r.actiondb[ACTION_PRE] = make([]string, 0)
	r.actiondb[ACTION_POST] = make([]string, 0)
	name := mw.Name()
	_, exists := r.db[name]
	if exists {
		log.E("ignoring already registered: %s", name)
		return
	}
	for _, act := range actions {
		r.actiondb[act] = append(r.actiondb[act], name)
	}
	r.db[name] = mw
}

func (r *mwRegistry) Enable(settings []*Settings) error {
	r.enable[ACTION_PRE] = make([]string, 0)
	r.enable[ACTION_POST] = make([]string, 0)
	for _, s := range settings {
		if _, ok := r.db[s.Name]; !ok {
			return fmt.Errorf("invalid middleware: %s", s.Name)
		}
		for _, act := range r.mwActions(s.Name) {
			r.enable[act] = append(r.enable[act], s.Name)
		}
	}
	return nil
}

func (r *mwRegistry) mwActions(name string) []MiddlewareAction {
	l := make([]MiddlewareAction, 0)
	actl := []MiddlewareAction{ACTION_PRE, ACTION_POST}
	for _, act := range actl {
		for _, dbname := range r.actiondb[act] {
			if dbname == name {
				l = append(l, act)
				break
			}
		}
	}
	return l
}

func (r *mwRegistry) GetAll(action MiddlewareAction) []Middleware {
	l := make([]Middleware, 0)
	added := make(map[string]bool)
	for _, name := range r.enable[action] {
		_, ok := added[name]
		if ok {
			l = append(l, r.db[name])
			added[name] = true
		} else {
			log.E("ignoring duplicate '%s' for action '%s'",
				name, action)
		}
	}
	return l
}

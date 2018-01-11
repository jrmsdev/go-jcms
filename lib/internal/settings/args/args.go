package args

import (
	"fmt"
	"strings"
)

type Args struct {
	prefix string            `json:"omit"`
	Args   map[string]string `json:",omitempty"`
}

func (a *Args) SetPrefix(prefix string) {
	a.prefix = prefix
}

func (a *Args) Get(key, defval string) *Value {
	v, ok := a.Args[a.getKey(key)]
	if ok {
		return newValue(v)
	}
	return newValue(defval)
}

func (a *Args) GetAll(base string) map[string]*Value {
	all := make(map[string]*Value)
	for _, k := range a.listKeys(base) {
		v, ok := a.Args[k]
		if !ok {
			v = "ERROR:args.GetAll"
		}
		all[k] = newValue(v)
	}
	return all
}

func (a *Args) getKey(key string) string {
	if a.prefix != "" && a.prefix != "." {
		return fmt.Sprintf("%s.%s", a.prefix, key)
	}
	return fmt.Sprintf("%s", key)
}

func (a *Args) listKeys(base string) []string {
	l := make([]string, 0)
	kfilter := false
	kprefix := a.getKey(base)
	if kprefix != "" && kprefix != "." {
		kfilter = true
		kprefix = fmt.Sprintf("%s.", kprefix)
	}
	for k := range a.Args {
		if kfilter {
			if strings.HasPrefix(k, kprefix) {
				l = append(l, k)
			}
		} else {
			l = append(l, k)
		}
	}
	return l
}

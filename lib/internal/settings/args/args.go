package args

import (
	"fmt"
)

type Args struct {
	prefix string            `json:"omit"`
	Args   map[string]string `json:",omitempty"`
}

func (a *Args) SetPrefix(prefix string) {
	a.prefix = prefix
}

func (a *Args) Get(key, defval string) string {
	v, ok := a.Args[a.getKey(key)]
	if ok {
		return v
	}
	return defval
}

func (a *Args) getKey(key string) string {
	if a.prefix != "" && a.prefix != "." {
		return fmt.Sprintf("%s.%s", a.prefix, key)
	}
	return fmt.Sprintf("%s", key)
}

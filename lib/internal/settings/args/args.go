package args

type Args struct {
	prefix string            `json:"omit"`
	Args   map[string]string `json:",omitempty"`
}

func (a *Args) SetPrefix(prefix string) {
	a.prefix = prefix
}

func (a *Args) Get(key, defval string) string {
	v, ok := a.Args[key]
	if ok {
		return v
	}
	return defval
}

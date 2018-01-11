package args

type Args struct {
	Args map[string]string `json:",omitempty"`
}

func (a *Args) Get(key, defval string) string {
	v, ok := a.Args[key]
	if ok {
		return v
	}
	return defval
}

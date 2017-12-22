package text

import "github.com/jrmsdev/go-jcms/internal/doctype/base"

type Engine struct {
    base.Engine
}

func New () *Engine {
    return &Engine {base.New ("text")}
}

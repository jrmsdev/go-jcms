package base

import "fmt"

type Engine struct {
    dtype string
}

func New (dtype string) Engine {
    return Engine{dtype}
}

func (e *Engine) Type () string {
    return e.dtype
}

func (e *Engine) String () string {
    return fmt.Sprintf ("doctype.%s", e.dtype)
}

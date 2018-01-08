package doctype

import "context"

type Engine interface {
	Type() string
	String() string
	Handle(context.Context) context.Context
}

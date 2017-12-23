// public API
package jcms

import (
    "github.com/jrmsdev/go-jcms/internal/core"
)

func Listen () string {
    return core.Listen ()
}

func Serve () {
    core.Serve ()
}

func Stop () {
    core.Stop ()
}

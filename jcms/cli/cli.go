// cli public API
package cli

import (
    "github.com/jrmsdev/go-jcms/internal/cli"
)

func Main () {
    cli.Main ()
}

func Webview (uri string) {
    cli.Webview (uri)
}

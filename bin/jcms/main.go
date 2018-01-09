package main

import (
	"os"

	"github.com/jrmsdev/go-jcms/lib/jcms"
)

func main() {
	err := jcms.LogFile(os.Stderr)
	if err != nil {
		panic(err)
	}
	defer jcms.LogClose()
	jcms.Listen()
	jcms.Serve()
}

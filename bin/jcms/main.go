package main

import (
	"os"

	"github.com/jrmsdev/go-jcms/lib/jcms"
)

func main() {
	err := jcms.LogStart("error", os.Stderr)
	if err != nil {
		panic(err)
	}
	defer jcms.LogStop()
	jcms.Listen()
	jcms.Serve()
}

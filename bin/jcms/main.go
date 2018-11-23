package main

import (
	"fmt"
	"os"

	"github.com/jrmsdev/go-jcms/bin/internal/flags"
	"github.com/jrmsdev/go-jcms/lib/jcms"
)

func main() {
	flags.Parse()
	err := jcms.LogStart(flags.LogLevel, os.Stderr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer jcms.LogStop()
	jcms.Listen(flags.HttpAddr)
	defer jcms.Stop()
	jcms.Serve()
}

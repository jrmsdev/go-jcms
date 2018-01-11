package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jrmsdev/go-jcms/lib/jcms"
)

var loglevel string

func init() {
	flag.StringVar(&loglevel, "log", "error", "set log `level`")
}

func main() {
	flag.Parse()
	err := jcms.LogStart(loglevel, os.Stderr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer jcms.LogStop()
	jcms.Listen()
	jcms.Serve()
}

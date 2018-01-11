package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/jrmsdev/go-jcms/lib/jcms"
)

var log = jcms.Logger("jcms-devel")
var loglevel string

func init() {
	flag.StringVar(&loglevel, "log", "debug", "set log `level`")
}

func main() {
	// parse command args
	flag.Parse()
	uri := path.Clean(flag.Arg(0))
	if uri == "." || uri == "" {
		uri = "/"
	}
	// jcms log
	err := jcms.LogStart(loglevel, os.Stderr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer jcms.LogStop()
	// use devel webapp
	if err = os.Setenv("JCMS_WEBAPP", "devel"); err != nil {
		log.Panic(err.Error())
	}
	// force using GOPATH as basedir
	gp, ok := os.LookupEnv("GOPATH")
	gp = filepath.Clean(gp)
	if !ok || gp == "." {
		log.Panic("GOPATH is not set")
	}
	// TODO: support possible ':' separator in GOPATH
	gp = filepath.Join(gp, "src", "github.com", "jrmsdev", "go-jcms", "webapps")
	if err := os.Setenv("JCMS_BASEDIR", gp); err != nil {
		log.Panic(err.Error())
	}
	// launch webview
	Webview(uri)
}

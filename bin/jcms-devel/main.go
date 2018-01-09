package main

import (
	"flag"
	"os"
	"path"
	"path/filepath"

	"github.com/jrmsdev/go-jcms/lib/jcms"
)

var log = jcms.Logger("jcms-devel")

func main() {
	err := jcms.LogStart("verbose", os.Stderr)
	if err != nil {
		panic(err)
	}
	defer jcms.LogStop()
	// use devel webapp
	if err = os.Setenv("JCMS_WEBAPP", "devel"); err != nil {
		log.Panic(err.Error())
	}
	// force using GOPATH
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
	// parse command args
	flag.Parse()
	uri := path.Clean(flag.Arg(0))
	if uri == "." || uri == "" {
		uri = "/"
	}
	// launch webview
	Webview(uri)
}

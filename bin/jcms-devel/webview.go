package main

import (
	"log"
	"net/url"

	"github.com/jrmsdev/go-jcms/lib/jcms"
	xwv "github.com/zserge/webview"
)

const (
	webviewResize = true
	webviewWidth  = 800
	webviewHeight = 600
)

func Webview(req string) {
	uri, err := url.Parse(jcms.Listen())
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("webview: req", req)
	go func() {
		jcms.Serve()
	}()
	uri.Path = req
	log.Println("webview: open", uri.String())
	xwv.Open("jcms", uri.String(), webviewWidth, webviewHeight, webviewResize)
	jcms.Stop()
}

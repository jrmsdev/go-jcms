package main

import (
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
		log.Panic(err.Error())
	}
	log.D("webview: req %#v", req)
	go func() {
		jcms.Serve()
	}()
	uri.Path = req
	log.D("webview: open %s", uri.String())
	xwv.Open("jcms", uri.String(), webviewWidth, webviewHeight, webviewResize)
	jcms.Stop()
}

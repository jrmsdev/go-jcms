package httpd

import (
	"context"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/jrmsdev/go-jcms/lib/internal/logger"
)

var log = logger.New("httpd")
var addr = "127.0.0.1:0"
var servemux = http.NewServeMux()
var listener net.Listener

var server = &http.Server{
	Addr:           addr,
	Handler:        servemux,
	ReadTimeout:    10 * time.Second,
	WriteTimeout:   10 * time.Second,
	MaxHeaderBytes: 1 << 20,
}

func Listen() *url.URL {
	var err error
	listener, err = net.Listen("tcp4", addr)
	if err != nil {
		log.Panic(err.Error())
	}
	url := &url.URL{}
	url.Scheme = "http"
	url.Host = listener.Addr().String()
	url.Path = "/"
	return url
}

func Serve() {
	log.D("serve")
	if listener == nil {
		log.Panic("nil listener... call httpd.Listen() first")
	}
	var err error
	err = server.Serve(listener)
	if err != nil {
		log.Panic(err.Error())
	}
}

func Stop() {
	log.D("stop")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		log.E("shutdown:", err)
	}
}

func Handle(prefix string, handler http.Handler) {
	servemux.Handle(prefix, handler)
}

func HandleFunc(prefix string, fn func(http.ResponseWriter, *http.Request)) {
	servemux.HandleFunc(prefix, fn)
}

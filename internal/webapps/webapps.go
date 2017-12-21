package webapps

import (
    "net/http"
    "github.com/jrmsdev/go-jcms/internal/httpd"
)

func init() {
    httpd.HandleFunc ("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("<html><body><p>welcome to jcms!</p></body></html>"))
    })
}

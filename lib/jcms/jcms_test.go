package jcms

import (
    "testing"
    "net/http"
)

func TestServe (t *testing.T) {
    uri := Listen ()
    go func(){
        Serve ()
    }()
    t.Log (uri)
    resp, err := http.Get (uri)
    if err != nil {
        t.Fatal (err)
    }
    resp.Body.Close ()
}

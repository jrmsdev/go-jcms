package main

import (
    "github.com/jrmsdev/go-jcms/jcms"
)

func main () {
    uri := jcms.Listen ()
    println (uri)
    jcms.Serve ()
}

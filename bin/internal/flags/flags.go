package flags

import "flag"

var LogLevel string
var HttpAddr string

func init() {
	flag.StringVar(&LogLevel, "log", "error", "set log `level`")
	flag.StringVar(&HttpAddr, "http", "127.0.0.1:0", "bind to http `address`")
}

func Parse () {
	flag.Parse()
}

func Arg (n int) string {
	return flag.Arg(n)
}

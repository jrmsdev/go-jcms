package args

import (
	"testing"
)

func TestArgs(t *testing.T) {
	a := &Args{}
	a.Args = map[string]string{
		"arg0":                "val0",
		"arg1":                "val1",
		"testing.prefix.arg0": "val0",
	}
	testGet(t, a)
	testPrefix(t, a)
}

func testGet(t *testing.T, a *Args) {
	v := a.Get("arg0", "GETERROR")
	if v.String() != "val0" {
		t.Error("invalid arg value for arg0:", v)
	}
	v = a.Get("arg", "DEFVAL")
	if v.String() != "DEFVAL" {
		t.Error("invalid arg default value:", v)
	}
}

func testPrefix(t *testing.T, a *Args) {
	a.SetPrefix("testing.prefix")
	v := a.Get("arg0", "GETERROR")
	if v.String() != "val0" {
		t.Error("invalid arg value for arg0:", v)
	}
	v = a.Get("arg", "DEFVAL")
	if v.String() != "DEFVAL" {
		t.Error("invalid arg default value:", v)
	}
	v = a.Get("arg1", "NOTFOUND")
	if v.String() != "NOTFOUND" {
		t.Error("prefixed arg1 key should be not found:", v)
	}
	a.SetPrefix("")
}

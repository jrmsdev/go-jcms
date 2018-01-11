package args

import (
	"testing"
)

func TestArgs(t *testing.T) {
	a := &Args{}
	a.Args = map[string]string{
		"arg0":                     "val0",
		"arg1":                     "val1",
		"testing.prefix.arg0":      "val0",
		"testing.value.int":        "128",
		"testing.value.int64":      "128",
		"testing.value.float":      "128.9",
		"testing.value.bool.true":  "true",
		"testing.value.bool.false": "false",
	}
	testGet(t, a)
	testPrefix(t, a)
	testValue(t, a)
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

func testValue(t *testing.T, a *Args) {
	var err error
	a.SetPrefix("testing.value")
	var intval int
	v := a.Get("int", "ERROR")
	intval, err = v.Int()
	if err != nil {
		t.Error(err)
	}
	if intval != 128 {
		t.Error("invalid value int:", intval)
	}
	var int64val int64
	v = a.Get("int64", "ERROR")
	int64val, err = v.Int64()
	if err != nil {
		t.Error(err)
	}
	if int64val != 128 {
		t.Error("invalid value int64:", int64val)
	}
	var floatval float64
	v = a.Get("float", "ERROR")
	floatval, err = v.Float()
	if err != nil {
		t.Error(err)
	}
	if floatval != 128.9 {
		t.Error("invalid value float:", floatval)
	}
	var boolval bool
	v = a.Get("bool.true", "ERROR")
	boolval, err = v.Bool()
	if err != nil {
		t.Error(err)
	}
	if boolval != true {
		t.Error("invalid value bool.true:", boolval)
	}
	v = a.Get("bool.false", "ERROR")
	boolval, err = v.Bool()
	if err != nil {
		t.Error(err)
	}
	if boolval != false {
		t.Error("invalid value bool.false:", boolval)
	}
}

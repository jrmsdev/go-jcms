package views

import (
	"testing"
)

func TestView(t *testing.T) {
	v := &View{
		Name:     "fakeview",
		Path:     "/path",
		Doctype:  "dtype",
		UseView:  "",
		Redirect: "",
		Location: "",
	}
	vs := v.String()
	if vs != "view:fakeview" {
		t.Error("invalid view String:", vs)
	}
}

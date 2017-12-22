package static

import (
    "testing"
    "github.com/jrmsdev/go-jcms/internal/doctype"
)

func TestEngine (t *testing.T) {
    e := New ()
    t.Log (e)
    testType (t, e)
}

func testType (t *testing.T, e doctype.Engine) {
    if e.Type () != "static" {
        t.Error (".Type != static")
    }
}

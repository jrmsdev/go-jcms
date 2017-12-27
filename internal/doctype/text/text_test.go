package text

import (
    "testing"
    "github.com/jrmsdev/go-jcms/internal/doctype"
)

func TestEngine (t *testing.T) {
    e, err := doctype.GetEngine ("text")
    if err != nil {
        t.Fatal (err)
    }
    t.Log (e)
    testType (t, e)
}

func testType (t *testing.T, e doctype.Engine) {
    if e.Type () != "text" {
        t.Error (".Type != text")
    }
}

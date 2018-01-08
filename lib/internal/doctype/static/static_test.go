package static

import (
	"testing"

	"github.com/jrmsdev/go-jcms/lib/internal/doctype"
)

func TestEngine(t *testing.T) {
	e, err := doctype.GetEngine("static")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(e)
	testType(t, e)
}

func testType(t *testing.T, e doctype.Engine) {
	if e.Type() != "static" {
		t.Error(".Type != static")
	}
}

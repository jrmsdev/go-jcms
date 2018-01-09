package base

import "testing"

func TestEngineBase(t *testing.T) {
	e := New("testengine")
	et := e.Type()
	if et != "testengine" {
		t.Error("invalid engine type:", et)
	}
	es := e.String()
	if es != "doctype.testengine" {
		t.Error("invalid engine string:", es)
	}
}

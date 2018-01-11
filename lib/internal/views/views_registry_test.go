package views

import (
	"testing"

	"github.com/jrmsdev/go-jcms/lib/internal/settings/view"
)

var testViews []*view.Settings

func init() {
	testViews = make([]*view.Settings, 0)
	testViews = append(testViews, &view.Settings{
		Name: "home",
		Path: "/",
	})
	testViews = append(testViews, &view.Settings{
		Name: "view0",
		Path: "/pathto/view0",
	})
	testViews = append(testViews, &view.Settings{
		Name: "view1",
		Path: "/pathto/view1",
	})
	testViews = append(testViews, &view.Settings{
		Name:    "view2",
		Path:    "/pathto/view2",
		UseView: "view1",
	})
	testViews = append(testViews, &view.Settings{
		Name:    "view3",
		Path:    "/pathto/view3",
		UseView: "notexistentview",
	})
}

func TestRegistry(t *testing.T) {
	reg := Register(testViews)
	testGet(t, reg)
	testGetNotFound(t, reg)
	testUseView(t, reg)
	testUseViewNotFound(t, reg)
}

func testGet(t *testing.T, reg *Registry) {
	v, err := reg.Get("/pathto/view0")
	if err != nil {
		t.Error("get view should not have failed:", err)
	}
	if v.Name != "view0" {
		t.Error("got invalid view:", v)
	}
}

func testGetNotFound(t *testing.T, reg *Registry) {
	v, err := reg.Get("/pathto/notexistent/view")
	if err == nil {
		t.Error("get view should have failed")
		return
	}
	if v != nil {
		t.Error("got a not nil view:", v)
		return
	}
}

func testUseView(t *testing.T, reg *Registry) {
	v, err := reg.Get("/pathto/view2")
	if err != nil {
		t.Error("get view should not have failed:", err)
	}
	if v.Name != "view1" {
		t.Error("got invalid useview:", v)
	}
}

func testUseViewNotFound(t *testing.T, reg *Registry) {
	v, err := reg.Get("/pathto/view3")
	if err == nil {
		t.Error("get view should have failed")
		return
	}
	if v != nil {
		t.Error("got a not nil view:", v)
		return
	}
}

package app

import (
	"encoding/xml"
	"io/ioutil"

	"github.com/jrmsdev/go-jcms/lib/internal/middleware"
	"github.com/jrmsdev/go-jcms/lib/internal/views"
)

type Settings struct {
	XMLName    xml.Name               `xml:"webapp"`
	Views      []*views.View          `xml:"view"`
	Middleware []*middleware.Settings `xml:"middleware"`
}

func newSettings(blob []byte) (*Settings, error) {
	s := &Settings{}
	err := xml.Unmarshal(blob, s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func readSettings(fn string) (*Settings, error) {
	blob, err := ioutil.ReadFile(fn)
	if err != nil {
		return nil, err
	}
	return newSettings(blob)
}

package views

import (
	"encoding/xml"
	"fmt"
)

type View struct {
	XMLName  xml.Name `xml:"view"`
	Name     string   `xml:"name,attr"`
	Path     string   `xml:"path,attr"`
	Doctype  string   `xml:"doctype,attr"`
	UseView  string   `xml:"useview,attr"`
	Redirect string   `xml:"redirect,attr"`
	Location string   `xml:"location,attr"`
}

func (v *View) String() string {
	return fmt.Sprintf("view:%s", v.Name)
}

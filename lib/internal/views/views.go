package views

import (
    "fmt"
    "encoding/xml"
    // init doctype engines
    _ "github.com/jrmsdev/go-jcms/lib/internal/doctype/base/loader"
)

type View struct {
    XMLName xml.Name `xml:"view"`
    Name string `xml:"name,attr"`
    Path string `xml:"path,attr"`
    IsHome string `xml:"home,attr"`
    Doctype string `xml:"doctype,attr"`
}

func (v *View) String () string {
    return fmt.Sprintf ("<view:%s>", v.Name)
}

package views

import (
    "encoding/xml"
)

type View struct {
    XMLName xml.Name `xml:"view"`
    Name string `xml:"name,attr"`
    Path string `xml:"path,attr"`
    IsHome string `xml:"home,attr"`
}

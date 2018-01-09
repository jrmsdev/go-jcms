package middleware

import (
	"encoding/xml"
	"fmt"
)

type Settings struct {
	XMLName xml.Name `xml:"middleware"`
	Name    string   `xml:"name,attr"`
}

func (s *Settings) String() string {
	return fmt.Sprintf("middleware.settings:%s", s.Name)
}

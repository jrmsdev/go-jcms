package app

import (
    "testing"
    "encoding/xml"
)

func TestSettings (t *testing.T) {
    data := `<?xml version="1.0" encoding="UTF-8"?>
             <webapp><invalid></invalid></webapp>`
    s, err := newSettings ([]byte(data))
    if err != nil {
        t.Fatal (err)
    }
    testXMLName (t, s.XMLName)
    testMarshalOutput (t, s)
}

func testXMLName (t *testing.T, xn xml.Name) {
    t.Logf ("%#v", xn)
    if xn.Space != "" {
        t.Error (".Space should be an empty string")
    }
    if xn.Local != "webapp" {
        t.Error (".Local != webapp")
    }
}

func testMarshalOutput (t *testing.T, s *Settings) {
    out, err := xml.Marshal (s)
    if err != nil {
        t.Fatal (err)
    }
    t.Log (string(out))
    if string(out) != "<webapp></webapp>" {
        t.Error ("invalid settings marshal")
    }
}

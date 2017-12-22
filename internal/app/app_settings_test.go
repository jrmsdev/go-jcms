package app

import (
    "testing"
    "encoding/xml"
)

func TestSettings (t *testing.T) {
    data := `<?xml version="1.0" encoding="UTF-8"?>
             <webapp><invalid></invalid></webapp>`
    s := &Settings{}
    err := xml.Unmarshal ([]byte(data), s)
    if err != nil {
        t.Fatal (err)
    }
    t.Logf ("%#v", s.XMLName)
    if s.XMLName.Space != "" {
        t.Error ("XMLName.Space should be an empty string")
    }
    if s.XMLName.Local != "webapp" {
        t.Error ("XMLName.Local != webapp")
    }
    out, err2 := xml.Marshal (s)
    if err2 != nil {
        t.Fatal (err)
    }
    t.Log (string(out))
    if string(out) != "<webapp></webapp>" {
        t.Error ("invalid settings marshal")
    }
}

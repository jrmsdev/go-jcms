package app

import (
    "io/ioutil"
    "encoding/xml"
)

type Settings struct {
    XMLName xml.Name `xml:"webapp"`
}

func newSettings (blob []byte) (*Settings, error) {
    s := &Settings{}
    err := xml.Unmarshal (blob, s)
    if err != nil {
        return nil, err
    }
    return s, nil
}

func readSettings (fn string) (*Settings, error) {
    blob, err := ioutil.ReadFile (fn)
    if err != nil {
        return nil, err
    }
    return newSettings (blob)
}

package app

import (
	"encoding/xml"
	"testing"

	"github.com/jrmsdev/go-jcms/lib/internal/env"
)

func TestNewSettings(t *testing.T) {
	data := `<?xml version="1.0" encoding="UTF-8"?>
             <webapp><invalid></invalid></webapp>`
	s, err := newSettings([]byte(data))
	if err != nil {
		t.Fatal(err)
	}
	testXMLName(t, s.XMLName, "webapp")
	testMarshalOutput(t, s)
}

func testXMLName(t *testing.T, xn xml.Name, val string) {
	t.Logf("%#v", xn)
	if xn.Space != "" {
		t.Error(".Space should be an empty string")
	}
	if xn.Local != val {
		t.Error(".Local !=", val)
	}
}

func testMarshalOutput(t *testing.T, s *Settings) {
	out, err := xml.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(out))
	if string(out) != "<webapp></webapp>" {
		t.Error("invalid settings marshal")
	}
}

func TestNewSettingsError(t *testing.T) {
	data := `<?xml version="1.0" encoding="UTF-8"?>
             <webapp><invalid></invalid></`
	_, err := newSettings([]byte(data))
	if err == nil {
		t.Error("xml unmarshal should have failed")
	}
}

func TestReadSettings(t *testing.T) {
	testappEnv("testing")
	defer testappEnv("") // cleanup
	fn := env.SettingsFile()
	t.Log(fn)
	s, err := readSettings(fn)
	if err != nil {
		t.Fatal(err)
	}
	testXMLName(t, s.XMLName, "webapp")
	testDevelViews(t, s)
}

func testDevelViews(t *testing.T, s *Settings) {
	for _, v := range s.Views {
		testXMLName(t, v.XMLName, "view")
		t.Log(v)
	}
}

func TestReadSettingsError(t *testing.T) {
	_, err := readSettings("/invalid/file/name")
	if err == nil {
		t.Error("ioutil.Readfile should have failed")
	}
}

package settings

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/jrmsdev/go-jcms/lib/internal/settings/middleware"
	"github.com/jrmsdev/go-jcms/lib/internal/settings/view"
)

type Settings struct {
	ViewList       []*view.Settings       `json:"View"`
	MiddlewareList []*middleware.Settings `json:"Middleware"`
}

func New(filename string) (*Settings, error) {
	s := &Settings{}
	if err := s.readFile(filename); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Settings) readFile(filename string) error {
	buf, err := ioutil.ReadFile(filename)
	if err != nil && err != io.EOF {
		return err
	}
	err = json.Unmarshal(buf, &s)
	if err != nil {
		return err
	}
	return nil
}

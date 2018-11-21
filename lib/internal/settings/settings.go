package settings

import (
	"encoding/json"
	"io"

	"github.com/jrmsdev/go-jcms/lib/internal/asset"
	//~ "github.com/jrmsdev/go-jcms/lib/internal/logger"
	"github.com/jrmsdev/go-jcms/lib/internal/settings/middleware"
	"github.com/jrmsdev/go-jcms/lib/internal/settings/view"
)

//~ var log = logger.New("settings")

type Settings struct {
	Title          string
	Theme          string
	ViewList       []*view.Settings       `json:"View"`
	MiddlewareList []*middleware.Settings `json:"Middleware"`
}

func New() (*Settings, error) {
	s := &Settings{}
	if err := readFile(s); err != nil {
		return nil, err
	}
	return s, nil
}

func readFile(s *Settings) error {
	buf, err := asset.ReadFile("settings.json")
	if err != nil && err != io.EOF {
		return err
	}
	err = json.Unmarshal(buf, &s)
	if err != nil {
		return err
	}
	return nil
}

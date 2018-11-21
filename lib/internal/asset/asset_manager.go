package asset

import (
	"io/ioutil"

	"github.com/jrmsdev/go-jcms/lib/internal/env"
)

var manager Manager

func init() {
	manager = newManager()
}

type assetManager struct{}

func newManager () *assetManager {
	return &assetManager{}
}

func (m *assetManager) ReadFile(name string) ([]byte, error) {
	return ioutil.ReadFile(env.WebappFile(name))
}

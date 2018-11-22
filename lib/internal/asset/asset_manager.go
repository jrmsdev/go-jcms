package asset

import (
	"path/filepath"
	"io/ioutil"

	"github.com/jrmsdev/go-jcms/lib/internal/env"
)

var manager Manager

type assetManager struct{}

func newManager() *assetManager {
	return &assetManager{}
}

func checkManager() {
	if manager == nil {
		manager = newManager()
	}
}

func SetManager(newmanager Manager) {
	if manager != nil {
		log.Panic("asset manager already initialized")
	} else {
		manager = newmanager
	}
}

func (m *assetManager) ReadFile(parts ...string) ([]byte, error) {
	fn := env.WebappFile(filepath.Join(parts...))
	log.D("ioutil ReadFile %s", fn)
	return ioutil.ReadFile(fn)
}

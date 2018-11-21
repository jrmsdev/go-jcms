package asset

import (
	"io/ioutil"

	"github.com/jrmsdev/go-jcms/lib/internal/env"
	"github.com/jrmsdev/go-jcms/lib/internal/logger"
)

var log = logger.New("asset_manager")

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

func (m *assetManager) ReadFile(name string) ([]byte, error) {
	fn := env.WebappFile(name)
	log.D("ReadFile %s", fn)
	return ioutil.ReadFile(fn)
}

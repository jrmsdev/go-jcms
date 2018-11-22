package asset

import (
	"os"
	"io/ioutil"

	"github.com/jrmsdev/go-jcms/lib/jcms/api"
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

func (m *assetManager) Open(filename string) (api.AssetFile, error) {
	fn := env.WebappFile(filename)
	log.D("os Open %s", fn)
	return os.Open(fn)
}

func (m *assetManager) ReadFile(name string) ([]byte, error) {
	fn := env.WebappFile(name)
	log.D("ioutil ReadFile %s", fn)
	return ioutil.ReadFile(fn)
}

func (m *assetManager) Stat(filename string) (os.FileInfo, error) {
	fn := env.WebappFile(filename)
	return os.Stat(fn)
}

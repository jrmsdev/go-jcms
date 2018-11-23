package asset

import (
	"io/ioutil"
	"os"

	"github.com/jrmsdev/go-jcms/lib/internal/env"
	"github.com/jrmsdev/go-jcms/lib/jcms/api"
)

var manager api.AssetManager

type assetManager struct{}

func newManager() *assetManager {
	return &assetManager{}
}

func checkManager() {
	if manager == nil {
		manager = newManager()
	}
}

func SetManager(newmanager api.AssetManager) {
	if manager != nil {
		log.Panic("asset manager already initialized")
	} else {
		manager = newmanager
	}
}

func (m *assetManager) Open(filename string) (api.AssetFile, error) {
	fn := env.WebappFile(filename)
	log.D("os Open %s", fn)
	fi, err := os.Stat(fn)
	if err != nil {
		return nil, err
	}
	if fi.IsDir() {
		return nil, &os.PathError{
			Op:   "open",
			Path: fn,
			Err:  os.ErrNotExist,
		}
	}
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

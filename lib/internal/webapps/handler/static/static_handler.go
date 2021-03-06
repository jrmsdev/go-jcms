package static

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"

	"github.com/jrmsdev/go-jcms/lib/internal/app"
	"github.com/jrmsdev/go-jcms/lib/internal/asset"
	"github.com/jrmsdev/go-jcms/lib/internal/httpd"
	"github.com/jrmsdev/go-jcms/lib/internal/logger"
	"github.com/jrmsdev/go-jcms/lib/jcms/api"
)

var log = logger.New("static.handler")

func Handle(a *app.App) {
	log.D("app %s", a)
	httpd.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(fileSystem{"static"})))
}

type fileSystem struct {
	name string
}

func (fs fileSystem) Open(name string) (http.File, error) {
	fn := filepath.Join(fs.name, name)
	fh, err := asset.Open(fn)
	if err != nil {
		log.E("Open %s", err)
		return nil, err
	}
	return newFile(fn, fh), nil
}

type file struct {
	api.AssetFile
	name string
}

func newFile(fn string, fh api.AssetFile) http.File {
	log.D("newFile %s", fn)
	return &file{fh, fn}
}

func (f *file) Readdir(count int) ([]os.FileInfo, error) {
	log.D("Readdir %s", f.name)
	return nil, errors.New("dir not found")
}

func (f *file) Stat() (os.FileInfo, error) {
	log.D("Stat %s", f.name)
	return asset.Stat(f.name)
}

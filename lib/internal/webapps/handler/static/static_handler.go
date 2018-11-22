package static

import (
	"os"
	"net/http"
	"path/filepath"
	//~ "io/ioutil"

	"github.com/jrmsdev/go-jcms/lib/internal/asset"
	"github.com/jrmsdev/go-jcms/lib/internal/app"
	"github.com/jrmsdev/go-jcms/lib/internal/httpd"
	"github.com/jrmsdev/go-jcms/lib/internal/logger"
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
	asset.File
	name string
}

func newFile(fn string, fh asset.File) http.File {
	log.D("newFile %s", fn)
	return &file{fh, fn}
}

func (f *file) Readdir(count int) ([]os.FileInfo, error) {
	log.D("Readdir %s", f.name)
	log.Panic("static handler Readdir not implemented")
	return nil, nil
}

func (f *file) Stat() (os.FileInfo, error) {
	log.D("Stat %s", f.name)
	return asset.Stat(f.name)
}

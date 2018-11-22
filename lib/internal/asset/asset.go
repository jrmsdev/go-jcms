package asset

import (
	"os"
	"io"
	"path/filepath"

	"github.com/jrmsdev/go-jcms/lib/internal/logger"
)

var log = logger.New("asset")

type File interface {
	io.ReadSeeker
	io.Closer
}

type Manager interface {
	Open(filename string) (File, error)
	Stat(filename string) (os.FileInfo, error)
	ReadFile(name string) ([]byte, error)
}

func ReadFile(parts ...string) ([]byte, error) {
	fn := filepath.Join(parts...)
	log.D("ReadFile %s", fn)
	checkManager()
	return manager.ReadFile(fn)
}

func Open(parts ...string) (File, error) {
	fn := filepath.Join(parts...)
	log.D("Open %s", fn)
	checkManager()
	return manager.Open(fn)
}

func Stat(parts ...string) (os.FileInfo, error) {
	fn := filepath.Join(parts...)
	log.D("Stat %s", fn)
	checkManager()
	return manager.Stat(fn)
}

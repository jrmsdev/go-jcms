package asset

import (
	"io"

	"github.com/jrmsdev/go-jcms/lib/internal/logger"
)

var log = logger.New("asset")

type File interface {
	io.ReadSeeker
	io.Closer
}

type Manager interface {
	ReadFile(parts ...string) ([]byte, error)
}

func ReadFile(parts ...string) ([]byte, error) {
	log.D("ReadFile %#v", parts)
	checkManager()
	return manager.ReadFile(parts...)
}

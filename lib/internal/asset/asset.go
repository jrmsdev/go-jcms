package asset

import (
	"io"
)

type File interface {
	io.ReadSeeker
	io.Closer
}

type Manager interface {
	ReadFile(name string) ([]byte, error)
}

func ReadFile(name string) ([]byte, error) {
	return manager.ReadFile(name)
}

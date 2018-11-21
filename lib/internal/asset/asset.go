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
	checkManager()
	return manager.ReadFile(name)
}

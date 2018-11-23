package api

import (
	"io"
	"os"
)

type AssetFile interface {
	io.ReadSeeker
	io.Closer
}

type AssetManager interface {
	Open(filename string) (AssetFile, error)
	Stat(filename string) (os.FileInfo, error)
	ReadFile(name string) ([]byte, error)
}

package api

import (
	"io"
)

type AssetFile interface {
	io.ReadSeeker
	io.Closer
}

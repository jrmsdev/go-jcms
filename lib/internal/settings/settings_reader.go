package settings

import (
	"github.com/jrmsdev/go-jcms/lib/internal/settings/view"
)

type Reader struct {
	View *view.Settings
}

func NewReader(src *Settings) (*Reader, error) {
	return nil, nil
}

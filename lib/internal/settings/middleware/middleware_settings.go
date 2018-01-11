package middleware

import (
	"fmt"

	"github.com/jrmsdev/go-jcms/lib/internal/settings/args"
)

type Settings struct {
	Name    string
	args.Args
}

func (s *Settings) String() string {
	return fmt.Sprintf("middleware.settings:%s", s.Name)
}

package middleware

import (
	"fmt"

	"github.com/jrmsdev/go-jcms/lib/internal/settings/args"
)

type Settings struct {
	Name string
	args.Args
}

func (s Settings) ID() string {
	return fmt.Sprintf("middleware.%s", s.Name)
}

func (s *Settings) String() string {
	return fmt.Sprintf("%s: settings", s.ID())
}

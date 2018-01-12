package memdb

import (
	dbm "github.com/jrmsdev/go-jcms/lib/internal/db"
)

func init() {
	dbm.Register()
}

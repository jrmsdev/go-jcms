package memdb

import (
	dbm "github.com/jrmsdev/go-jcms/lib/internal/db"
	//~ "github.com/jrmsdev/go-jcms/lib/internal/db/query"
	"github.com/jrmsdev/go-jcms/lib/internal/logger"
)

const jcmsid = "db.memdb"

var log = logger.New(jcmsid)

func Init() {
	dbm.Register(jcmsid, newEngine())
}

type Engine struct{}

func newEngine() *Engine {
	return &Engine{}
}

func (e *Engine) Name() string {
	return jcmsid
}

func (e *Engine) Open() {}

func (e *Engine) Close() {}

func (e *Engine) Set() {}

func (e *Engine) SetAll() {}

func (e *Engine) Get() {}

func (e *Engine) GetAll() {}

func (e *Engine) Del() {}

func (e *Engine) DelAll() {}

func (e *Engine) Update() {}

func (e *Engine) UpdateAll() {}

package db

import (
	"github.com/jrmsdev/go-jcms/lib/internal/logger"
)

var log = logger.New("db")

type Engine interface {
	Name() string
	Open()
	Close()
	Set()
	SetAll()
	Del()
	DelAll()
	Get()
	GetAll()
	Update()
	UpdateAll()
}

func Open(uri string) (Engine, error) {
	log.D("open %s", uri)
	eng, err := getEngine(uri)
	if err != nil {
		return nil, err
	}
	return eng, nil
}

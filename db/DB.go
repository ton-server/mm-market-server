package db

import (
	"github.com/sunjiangjun/xlog"
	"gorm.io/gorm"
)

type DB struct {
	core *gorm.DB
	log  *xlog.XLog
}

func NewDB(db *gorm.DB, log *xlog.XLog) *DB {
	return &DB{
		core: db,
		log:  log,
	}
}

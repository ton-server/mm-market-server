package server

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/sunjiangjun/xlog"
	"github.com/ton-server/mm-market-server/common/driver"
	"github.com/ton-server/mm-market-server/config"
	"github.com/ton-server/mm-market-server/db"
)

type Monitor struct {
	db  *db.DB
	log *logrus.Entry
	ctx context.Context
}

func NewMonitor(cfg *config.DB, log *xlog.XLog, ctx context.Context) *Monitor {
	conn, err := driver.Open(cfg.User, cfg.Password, cfg.Addr, cfg.DbName, cfg.Port, log)
	if err != nil {
		panic(err)
	}

	pg := db.NewDB(conn, log)

	return &Monitor{
		db:  pg,
		ctx: ctx,
		log: log.WithField("module", "handler"),
	}
}

func (m *Monitor) Start() {
	//for {
	//	select {
	//	case <-m.ctx.Done():
	//		return
	//	default:
	//		m.loop()
	//	}
	//}
}

func (m *Monitor) loop() {

}

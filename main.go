package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sunjiangjun/xlog"
	"github.com/ton-server/mm-market-server/config"
	"github.com/ton-server/mm-market-server/server"
)

func main() {

	var configPath string
	flag.StringVar(&configPath, "config", "./config.json", "The system file of config")
	flag.Parse()
	if len(configPath) < 1 {
		panic("can not find config file")
	}
	cfg := config.LoadConfig(configPath)

	if cfg.Log == nil {
		cfg.Log.Delay = 2
		cfg.Log.Path = "./log/market"
	}

	LOG := xlog.NewXLogger().BuildOutType(xlog.FILE).BuildLevel(xlog.InfoLevel).BuildFormatter(xlog.FORMAT_JSON).BuildFile(cfg.Log.Path, 24*time.Hour)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	monitor := server.NewMonitor(cfg.DB, LOG, ctx)
	go monitor.Start()

	gin.SetMode(gin.ReleaseMode)

	e := gin.New()
	root := e.Group(cfg.Root)
	root.Use(gin.LoggerWithConfig(gin.LoggerConfig{Output: LOG.Out}))

	srv := server.NewHandler(cfg.DB, LOG)
	root.GET("/coin/list", srv.GetCoinList)
	root.GET("/coin/info", srv.GetCoinInfo)
	root.GET("/coin/fullCoin", srv.GetCoin)
	root.POST("/history/create", srv.SubmitTxHistory)
	root.GET("/history/query", srv.GetTxHistory)
	root.POST("/user/create", srv.SubmitUser)
	root.GET("/user/query", srv.GetUser)
	root.POST("/user/update", srv.UpdateUser)

	err := e.Run(fmt.Sprintf(":%v", cfg.Port))
	if err != nil {
		panic(err)
	}

}

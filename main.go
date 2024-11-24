package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
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

	monitor := server.NewMonitor(cfg.DB, LOG, cfg.AdminAddress, cfg.TonHost, ctx)
	go monitor.Start()

	gin.SetMode(gin.ReleaseMode)

	e := gin.New()
	e.Use(CORSMiddleware())
	root := e.Group(cfg.Root)
	root.Use(gin.LoggerWithConfig(gin.LoggerConfig{Output: LOG.Out}))
	// 使用 CORS 中间件

	srv := server.NewHandler(cfg.DB, LOG)
	root.GET("/coin/list", srv.GetCoinList)
	root.GET("/coin/info", srv.GetCoinInfo)
	root.GET("/coin/fullCoin", srv.GetCoin)
	root.POST("coin/create", srv.SubmitCoin)
	root.POST("/history/create", srv.SubmitTxHistory)
	root.GET("/history/query", srv.GetTxHistory)
	root.GET("/user/queryOrCreate", srv.GetUser)

	err := e.Run(fmt.Sprintf(":%v", cfg.Port))
	if err != nil {
		panic(err)
	}

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // 设置允许的域名，`*` 表示允许所有
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true") // 允许凭据（cookies等）

		// 对于 OPTIONS 请求，直接返回 200
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}

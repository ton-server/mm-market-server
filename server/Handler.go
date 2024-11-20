package server

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/sunjiangjun/xlog"
	"github.com/tidwall/gjson"
	"github.com/ton-server/mm-market-server/common/driver"
	"github.com/ton-server/mm-market-server/config"
	"github.com/ton-server/mm-market-server/db"
)

const (
	TimeFormat = "2006-01-02 15:04:05"
)

type Handler struct {
	db  *db.DB
	log *logrus.Entry
}

func NewHandler(cfg *config.DB, log *xlog.XLog) *Handler {

	conn, err := driver.Open(cfg.User, cfg.Password, cfg.Addr, cfg.DbName, cfg.Port, log)
	if err != nil {
		panic(err)
	}

	pg := db.NewDB(conn, log)

	return &Handler{
		db:  pg,
		log: log.WithField("module", "handler"),
	}
}

func (h *Handler) Monitor(ctx *gin.Context) {
	b, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		h.Error(ctx, "", ctx.Request.RequestURI, err.Error())
		return
	}
	startTime := gjson.ParseBytes(b).Get("startTime").String()
	start, err := time.ParseInLocation(TimeFormat, startTime, time.UTC)
	if err != nil {
		h.Error(ctx, "", ctx.Request.RequestURI, err.Error())
		return
	}
	endTime := gjson.ParseBytes(b).Get("endTime").String()
	end, err := time.ParseInLocation(TimeFormat, endTime, time.UTC)
	if err != nil {
		h.Error(ctx, "", ctx.Request.RequestURI, err.Error())
		return
	}
	end = end.Add(-60 * time.Minute)
	if end.Before(start) {
		h.Error(ctx, "", ctx.Request.RequestURI, fmt.Errorf("end time is before start time").Error())
		return
	}
	//array := gjson.ParseBytes(b).Get("status").Array()
	//list := make([]int64, 0, 2)
	//for _, v := range array {
	//	list = append(list, v.Int())
	//}

	h.Success(ctx, string(b), nil, ctx.Request.RequestURI)
}

func (h *Handler) QueryTxs(ctx *gin.Context) {
	b, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		h.Error(ctx, "", ctx.Request.RequestURI, err.Error())
		return
	}

	mp := make(map[string]any, 2)

	h.Success(ctx, string(b), mp, ctx.Request.RequestURI)
}

func (h *Handler) Income(ctx *gin.Context) {
	b, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		h.Error(ctx, "", ctx.Request.RequestURI, err.Error())
		return
	}

	h.Success(ctx, string(b), nil, ctx.Request.RequestURI)
}

func (h *Handler) Pay(ctx *gin.Context) {
	b, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		h.Error(ctx, "", ctx.Request.RequestURI, err.Error())
		return
	}

	h.Success(ctx, string(b), nil, ctx.Request.RequestURI)
}

const (
	SUCCESS = 0
	FAIL    = 1
)

func (h *Handler) Success(c *gin.Context, req string, resp interface{}, path string) {
	req = strings.Replace(req, "\t", "", -1)
	req = strings.Replace(req, "\n", "", -1)
	if v, ok := resp.(string); ok {
		resp = strings.Replace(v, "\n", "", -1)
	}
	h.log.Printf("path=%v,req=%v,resp=%v\n", path, req, resp)
	mp := make(map[string]interface{})
	mp["code"] = SUCCESS
	mp["message"] = "ok"
	mp["data"] = resp
	c.JSON(200, mp)
}

func (h *Handler) Error(c *gin.Context, req string, path string, err string) {
	req = strings.Replace(req, "\t", "", -1)
	req = strings.Replace(req, "\n", "", -1)
	h.log.Errorf("path=%v,req=%v,err=%v\n", path, req, err)
	mp := make(map[string]interface{})
	mp["code"] = FAIL
	mp["message"] = err
	mp["data"] = ""
	c.JSON(200, mp)
}

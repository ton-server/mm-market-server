package server

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/sunjiangjun/xlog"
	"github.com/ton-server/mm-market-server/common/driver"
	"github.com/ton-server/mm-market-server/common/util"
	"github.com/ton-server/mm-market-server/config"
	"github.com/ton-server/mm-market-server/db"
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

func (h *Handler) GetCoinList(ctx *gin.Context) {
	currentPage := ctx.Query("currentPage")
	page, err := strconv.Atoi(currentPage)
	if err != nil {
		h.Error(ctx, "", ctx.Request.RequestURI, err.Error())
		return
	}

	pageSize := ctx.Query("pageSize")
	size, err := strconv.Atoi(pageSize)
	if err != nil {
		h.Error(ctx, "", ctx.Request.RequestURI, err.Error())
		return
	}

	fullCoin := ctx.Query("fullCoin")
	var full int = 0
	if len(fullCoin) > 0 {
		full, err = strconv.Atoi(fullCoin)
		if err != nil {
			h.Error(ctx, "", ctx.Request.RequestURI, err.Error())
			return
		}
	}

	list, total, err := h.db.GetCoinList(page, size, full == 1)
	if err != nil {
		h.Error(ctx, "", ctx.Request.RequestURI, err.Error())
		return
	}

	for _, v := range list {
		v.ExpireTime2 = v.ExpireTime.Unix()

		c, p, err := h.db.GetCoinPriceList(v.ContractAddress)
		if err != nil {
			continue
		}
		v.Usd = c.Price
		if p == nil {
			v.Change = "--"
		} else {
			ch, err := util.CalculatePercentageChange(p.Price, c.Price)
			if err != nil {
				continue
			}
			v.Change = fmt.Sprintf("%.2f%%", ch)
		}
	}

	mp := make(map[string]any, 2)
	mp["list"] = list
	mp["total"] = total

	h.Success(ctx, "", mp, ctx.Request.RequestURI)
}

func (h *Handler) GetCoin(ctx *gin.Context) {
	uuid := ctx.Query("uuid")
	r, err := h.db.GetCoinWithCoinInfo(uuid)
	if err != nil {
		h.Error(ctx, "", ctx.Request.RequestURI, err.Error())
		return
	}
	r.ExpireTime2 = r.ExpireTime.Unix()
	c, p, err := h.db.GetCoinPriceList(r.ContractAddress)
	if err != nil {
		h.Error(ctx, "", ctx.Request.RequestURI, err.Error())
		return
	}
	r.Usd = c.Price
	if p == nil {
		r.Change = "--"
	} else {
		ch, err := util.CalculatePercentageChange(p.Price, c.Price)
		if err != nil {
			h.Error(ctx, "", ctx.Request.RequestURI, err.Error())
			return
		}
		r.Change = fmt.Sprintf("%.2f%%", ch)
	}

	h.Success(ctx, "", r, ctx.Request.RequestURI)
}

func (h *Handler) GetCoinInfo(ctx *gin.Context) {
	uuid := ctx.Query("uuid")
	r, err := h.db.GetCoinInfo(uuid)
	if err != nil {
		h.Error(ctx, "", ctx.Request.RequestURI, err.Error())
		return
	}

	h.Success(ctx, "", r, ctx.Request.RequestURI)
}

func (h *Handler) GetTxHistory(ctx *gin.Context) {
	address := ctx.Query("address")
	list, err := h.db.GetTxHistoryByAddress(address)
	if err != nil {
		h.Error(ctx, "", ctx.Request.RequestURI, err.Error())
		return
	}

	for _, v := range list {
		v.CreateTime2 = v.CreateTime.Unix()
	}

	h.Success(ctx, "", list, ctx.Request.RequestURI)
}

func (h *Handler) SubmitTxHistory(ctx *gin.Context) {
	b, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		h.Error(ctx, "", ctx.Request.RequestURI, err.Error())
		return
	}

	var tx db.TxHistory
	err = json.Unmarshal(b, &tx)
	if err != nil {
		h.Error(ctx, string(b), ctx.Request.RequestURI, err.Error())
		return
	}

	err = h.db.NewTxHistory(&tx)
	if err != nil {
		h.Error(ctx, string(b), ctx.Request.RequestURI, err.Error())
		return
	}

	h.Success(ctx, string(b), nil, ctx.Request.RequestURI)
}

func (h *Handler) SubmitCoin(ctx *gin.Context) {
	b, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		h.Error(ctx, "", ctx.Request.RequestURI, err.Error())
		return
	}

	var rc db.RecommendCoin
	err = json.Unmarshal(b, &rc)
	if err != nil {
		h.Error(ctx, string(b), ctx.Request.RequestURI, err.Error())
		return
	}

	uid := uuid.NewString()
	rc.UUID = uid
	rc.CoinInfo.UUID = uid
	err = h.db.NewRecommendCoinAndCoinInfo(&rc)
	if err != nil {
		h.Error(ctx, string(b), ctx.Request.RequestURI, err.Error())
		return
	}

	h.Success(ctx, string(b), nil, ctx.Request.RequestURI)
}

func (h *Handler) SubmitUser(ctx *gin.Context) {
	b, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		h.Error(ctx, "", ctx.Request.RequestURI, err.Error())
		return
	}

	var u db.User
	err = json.Unmarshal(b, &u)
	if err != nil {
		h.Error(ctx, string(b), ctx.Request.RequestURI, err.Error())
		return
	}

	err = h.db.SubmitUser(&u)
	if err != nil {
		h.Error(ctx, string(b), ctx.Request.RequestURI, err.Error())
		return
	}

	h.Success(ctx, string(b), nil, ctx.Request.RequestURI)
}

func (h *Handler) GetUser(ctx *gin.Context) {
	address := ctx.Query("address")
	u, err := h.db.GetUser(address)
	if err != nil {
		h.Error(ctx, "", ctx.Request.RequestURI, err.Error())
		return
	}

	u.ExpireTime2 = u.ExpireTime.Unix()
	h.Success(ctx, "", u, ctx.Request.RequestURI)
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

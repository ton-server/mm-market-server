package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/sunjiangjun/xlog"
	"github.com/tidwall/gjson"
	"github.com/ton-server/mm-market-server/common/driver"
	"github.com/ton-server/mm-market-server/config"
	"github.com/ton-server/mm-market-server/db"
	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/ton"
)

type Monitor struct {
	db           *db.DB
	log          *logrus.Entry
	ctx          context.Context
	AdminAddress string
	TonHost      string
}

func NewMonitor(cfg *config.DB, log *xlog.XLog, AdminAddress, TonHost string, ctx context.Context) *Monitor {
	conn, err := driver.Open(cfg.User, cfg.Password, cfg.Addr, cfg.DbName, cfg.Port, log)
	if err != nil {
		panic(err)
	}

	pg := db.NewDB(conn, log)

	return &Monitor{
		db:           pg,
		ctx:          ctx,
		AdminAddress: AdminAddress,
		TonHost:      TonHost,
		log:          log.WithField("module", "handler"),
	}
}

func (m *Monitor) Start() {
	//for {
	//	<-time.After(20 * time.Second)
	//	select {
	//	case <-m.ctx.Done():
	//		return
	//	default:
	//		m.loop2()
	//	}
	//}
}

func (m *Monitor) loop2() {
	url := "%v/getTransactions?address=%v&limit=50&to_lt=0&archival=false"
	url = fmt.Sprintf(url, m.TonHost, m.AdminAddress)
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	root := gjson.ParseBytes(bs)
	if root.Get("ok").Bool() {

		//txs := root.Get("result").Array()

		//for _, tx := range txs {
		//hash := tx.Get("transaction_id.hash").String()
		//insource := tx.Get("in_msg.source").String()
		//indestination := tx.Get("in_msg.destination").String()
		//invalue := tx.Get("in_msg.value").String()
		//}

	}

}

func (m *Monitor) loop() {
	client := liteclient.NewConnectionPool()

	/*
		Mainnet public servers - https://ton.org/global.config.json
		Testnet public servers - https://ton-blockchain.github.io/testnet-global.config.json
	*/
	configUrl := "https://ton.org/global.config.json"
	err := client.AddConnectionsFromConfigUrl(context.Background(), configUrl)
	if err != nil {
		panic(err)
	}

	// initialize ton api lite connection wrapper
	api := ton.NewAPIClient(client, ton.ProofCheckPolicyFast).WithRetry()

	// if we want to route all requests to the same node, we can use it
	ctx := client.StickyContext(context.Background())

	// we need fresh block info to run get methods
	b, err := api.CurrentMasterchainInfo(ctx)
	if err != nil {
		log.Fatalln("get block err:", err.Error())
		return
	}

	// 查询交易记录

	// TON Foundation account
	addr := address.MustParseAddr(m.AdminAddress)

	// we use WaitForBlock to make sure block is ready,
	// it is optional but escapes us from liteserver block not ready errors
	account, err := api.WaitForBlock(b.SeqNo).GetAccount(ctx, b, addr)
	if err != nil {
		log.Fatalln("get account err:", err.Error())
		return
	}

	fmt.Printf("Is active: %v\n", account.IsActive)
	if account.IsActive {
		fmt.Printf("Status: %s\n", account.State.Status)
		fmt.Printf("Balance: %s TON\n", account.State.Balance.String())
		if account.Data != nil {
			fmt.Printf("Data: %s\n", account.Data.Dump())
		}
	}

	transactions, err := api.ListTransactions(context.Background(), addr, 50, account.LastTxLT, account.LastTxHash)
	if err != nil {
		log.Fatalf("Failed to get transactions: %v", err)
	}

	// 输出交易记录
	for _, tx := range transactions {
		fmt.Printf("LT: %v, Hash: %s, Value: %s TON\n", tx.LT, tx.Hash, tx.String())
	}
}

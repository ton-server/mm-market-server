package server

import (
	"context"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

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
		log:          log.WithField("module", "monitor"),
	}
}

func (m *Monitor) Start() {
	go m.stakeLoop()
	go m.coinPriceLoop()
}

// stakeLoop 监听会员充值
func (m *Monitor) coinPriceLoop() {
	for {
		<-time.After(20 * time.Minute)
		select {
		case <-m.ctx.Done():
			return
		default:
			m.priceLoop()
		}
	}
}

// stakeLoop 监听会员充值
func (m *Monitor) stakeLoop() {
	for {
		<-time.After(5 * time.Second)
		select {
		case <-m.ctx.Done():
			return
		default:
			m.loop2()
		}
	}
}

func (m *Monitor) loop2() {

	//a1, _ := address.ParseAddr("EQA3MJp-9cU-UlnXAWl1RT18wtcieZT9HvAYRxwYgSjTYSub")
	//a2, _ := address.ParseAddr("UQA3MJp-9cU-UlnXAWl1RT18wtcieZT9HvAYRxwYgSjTYXZe")
	//a1.Equals(a2)

	users, err := m.db.GetNormalUser()
	if err != nil {
		return
	}
	mp := make(map[string]*db.User, len(users))
	for _, v := range users {
		addr, err := address.ParseAddr(v.Address)
		if err != nil {
			continue
		}
		key := fmt.Sprintf("%v:%v", addr.Workchain(), hex.EncodeToString(addr.Data()))
		mp[key] = v
	}

	url := "%v/getTransactions?address=%v&limit=100&to_lt=0&archival=false"
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

		txs := root.Get("result").Array()

		for _, tx := range txs {
			if len(tx.Get("out_msgs").Array()) == 0 {
				hash := tx.Get("transaction_id.hash").String()
				insource := tx.Get("in_msg.source").String()
				a, err := address.ParseAddr(insource)
				if err != nil {
					continue
				}

				//indestination := tx.Get("in_msg.destination").String()
				invalue := tx.Get("in_msg.value").String()
				//m.log.Printf("workId:%v,addr:%v,sum:%v", a.Workchain(), base64.StdEncoding.EncodeToString(a.Data()), a.Checksum())
				//m.log.Printf("hash:%v,in.source:%v,in.destination:%v,in.value:%v \n", hash, insource, indestination, invalue)

				key := fmt.Sprintf("%v:%v", a.Workchain(), hex.EncodeToString(a.Data()))
				if u, ok := mp[key]; ok {
					_ = m.db.UpdateUser(u.Address, 1, hash, invalue, time.Now().UTC().Add(10000*24*time.Hour))
				}

			}
		}

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

func (m *Monitor) priceLoop() {

	list, err := m.db.GetActiveTask()
	if err != nil {
		return
	}

	for _, v := range list {
		addr := v.ContractAddress
		//r:=v.Rate
		price, err := GetPrice(addr)
		if err != nil {
			continue
		}

		err = m.db.NewCoinPrice(&db.CoinPriceRecord{
			ContractAddress: addr,
			Price:           price,
			RecordTime:      time.Now().UTC().Format(db.TimeFormat),
		})
		if err != nil {
			continue
		}

	}

}

func GetPrice(address string) (string, error) {
	//https://api.ston.fi/v1/assets/EQAvlWFDxGF2lXm67y4yzC17wYKD9A0guwPkMs1gOsM__NOT

	url := fmt.Sprintf("https://api.ston.fi/v1/assets/%v", address)
	resp, err := http.Get(url)
	if err != nil {
		return "", nil
	}

	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}

	defer resp.Body.Close()

	root := gjson.ParseBytes(bs)
	price := root.Get("asset.dex_usd_price").String()
	return price, nil
}

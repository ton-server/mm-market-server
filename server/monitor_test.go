package server

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"testing"

	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/ton"
)

func TestMonitor_loop(t *testing.T) {
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
	addr := address.MustParseAddr("UQA3MJp-9cU-UlnXAWl1RT18wtcieZT9HvAYRxwYgSjTYXZe")

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
		base64Hash := base64.StdEncoding.EncodeToString(tx.Hash)
		fmt.Printf("tx.hash:%v", base64Hash)

		//SenderAddr := tx.IO.In.Msg.SenderAddr()
		//DestAddr := tx.IO.In.Msg.DestAddr()
		fmt.Println(tx.String())

	}
}

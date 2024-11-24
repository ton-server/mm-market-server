package db

import (
	"fmt"
	"testing"
	"time"

	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/sunjiangjun/xlog"
	"github.com/ton-server/mm-market-server/common/driver"
)

func Init() *DB {
	x := xlog.NewXLogger()
	db, err := driver.Open("root", "123456789", "190.92.213.101", "ton-server", 3306, x)
	if err != nil {
		panic(err)
	}
	return NewDB(db, x)
}

func TestDB_GetCoinInfo(t *testing.T) {
	db := Init()
	r, err := db.GetCoinInfo("456e7890-f12b-34c5-d678-890123456789")
	assert.NoError(t, err)
	assert.Equal(t, "456e7890-f12b-34c5-d678-890123456789", r.UUID)
}

func TestDB_GetCoinList(t *testing.T) {
	db := Init()
	list, total, err := db.GetCoinList(1, 10, true)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), total)
	assert.Equal(t, list[0].UUID, "456e7890-f12b-34c5-d678-890123456789")
}

func TestDB_GetCoinWithCoinInfo(t *testing.T) {
	db := Init()
	r, err := db.GetCoinWithCoinInfo("456e7890-f12b-34c5-d678-890123456789")
	assert.NoError(t, err)
	assert.Equal(t, "456e7890-f12b-34c5-d678-890123456789", r.UUID)
	assert.Equal(t, "456e7890-f12b-34c5-d678-890123456789", r.CoinInfo.UUID)
}

func TestDB_GetTxHistoryByAddress(t *testing.T) {
	db := Init()
	list, err := db.GetTxHistoryByAddress("0xabcdef1234567890abcdef1234567890abcdef12")
	assert.NoError(t, err)
	assert.Equal(t, "0xabcdef1234567890abcdef1234567890abcdef12", list[0].FromAddress)
}

func TestDB_GetUser(t *testing.T) {
	db := Init()
	u, err := db.GetUser("0x1234567890abcdef2")
	assert.NoError(t, err)
	assert.Equal(t, "0x1234567890abcdef2", u.Address)
}

func TestDB_NewCoinInfo(t *testing.T) {
	coinInfo := CoinInfo{
		Id:     1,
		UUID:   "456e7890-f12b-34c5-d678-890123456789",
		Detail: "This is a detailed description of the coin.",
	}
	db := Init()
	err := db.NewCoinInfo(&coinInfo)
	assert.NoError(t, err)
}

func TestDB_NewRecommendCoin(t *testing.T) {

	uid := uuid.NewString()

	mp := make(map[string]any)
	mp["image"] = "https://cache.tonapi.io/imgproxy/8pvtFabCMkidoyjRHg38rHKpSTXkoL8R4urZOeKwq18/rs:fill:200:200:1/g:no/aHR0cHM6Ly9pLnBvc3RpbWcuY2MvTHM4d0Y0TVAvSU1HLTAwMjcucG5n.webp"
	mp["description"] = "Welcome to the world of meme coin 'Durov, bring back the wall!' - for those who remember and love this legendary meme! This token has an issuance of 19,840,000 coins. The token was created by the community solely for your enjoyment and entertainment. Join us to relish the jokes and adventures associated with this fantastic meme. There is no purpose here except to bring a smile to your face and evoke positive emotions!"
	mp["owner"] = "UQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAJKZ"
	mp["recommended reason"] = ""
	mp["recommendation index"] = 97
	bs, _ := json.Marshal(mp)

	// 构建 CoinInfo 对象实例
	coinInfo := &CoinInfo{
		UUID:   uid,
		Detail: string(bs),
	}

	// 构建 RecommendCoin 对象实例
	recommendCoin := RecommendCoin{
		UUID:            uid,
		Symbol:          "WALL",
		Decimals:        9,
		TotalSupply:     "19839790",
		ContractAddress: "EQDdCha_K-Z97lKl599O0GDAt0py2ZUuons4Wuf85tq6NXIO",
		Index:           12,
		CoinInfo:        coinInfo,
		ExpireTime:      time.Now().UTC().Add(365 * 24 * time.Hour), // 1 年后过期
	}

	b, _ := json.Marshal(recommendCoin)
	fmt.Println(string(b))

	db := Init()
	err := db.NewRecommendCoinAndCoinInfo(&recommendCoin)
	assert.NoError(t, err)
}

func TestDB_NewTxHistory(t *testing.T) {
	txHistory := TxHistory{
		Id:              1,
		FromAddress:     "0xabcdef1234567890abcdef1234567890abcdef12",
		ToAddress:       "0x1234567890abcdef1234567890abcdef12345678",
		ContractAddress: "0xcontract1234567890abcdef1234567890abcdef",
		Amount:          "1000",
		TxId:            "0xtransaction1234567890abcdef1234567890abcdef",
		TxStatus:        1, // 成功
		TxInfo:          "Transaction successful and confirmed.",
		CreateTime:      time.Now(),
		UpdateTime:      time.Now(),
	}

	b, _ := json.Marshal(txHistory)
	fmt.Println(string(b))

	db := Init()
	err := db.NewTxHistory(&txHistory)
	assert.NoError(t, err)

}

func TestDB_SubmitUser(t *testing.T) {
	user := User{
		Id:          1,
		Name:        "Alice",
		Address:     "0x1234567890abcdef",
		Role:        1, // VIP
		StakeTx:     "0xabcdef1234567890",
		StakeAmount: "1000",
		ExpireTime:  time.Now().Add(30 * 24 * time.Hour), // 30 天后过期
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
	}

	b, _ := json.Marshal(user)
	fmt.Println(string(b))

	db := Init()
	err := db.SubmitUser(&user)
	assert.NoError(t, err)
}

func TestDB_UpdateUser(t *testing.T) {
	db := Init()
	err := db.UpdateUser("0x1234567890abcdef2", 2, "0xabcdef1234567890", "10000", time.Now().UTC().Add(1000*24*time.Hour))
	assert.NoError(t, err)
}

package db

import (
	"github.com/sunjiangjun/xlog"
	"gorm.io/gorm"
)

type DB struct {
	core *gorm.DB
	log  *xlog.XLog
}

func (D *DB) SubmitUser(u *User) error {
	//TODO implement me
	panic("implement me")
}

func (D *DB) UpdateUser(address string, m map[string]any) error {
	//TODO implement me
	panic("implement me")
}

func (D *DB) GetUser(address string) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (D *DB) NewRecommendCoin(rc *RecommendCoin) error {
	//TODO implement me
	panic("implement me")
}

func (D *DB) GetCoinList(currentPage int, pageSize int) ([]*RecommendCoin, int64, error) {
	//TODO implement me
	panic("implement me")
}

func (D *DB) NewCoinInfo(ci *CoinInfo) error {
	//TODO implement me
	panic("implement me")
}

func (D *DB) GetCoinInfo(uuid string) (*CoinInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (D *DB) GetCoinWithCoinInfo(address string) (*RecommendCoin, error) {
	//TODO implement me
	panic("implement me")
}

func (D *DB) NewTxHistory(tx *TxHistory) error {
	//TODO implement me
	panic("implement me")
}

func (D *DB) GetTxHistoryByAddress(address string) (*TxHistory, error) {
	//TODO implement me
	panic("implement me")
}

func NewDB(db *gorm.DB, log *xlog.XLog) *DB {
	return &DB{
		core: db,
		log:  log,
	}
}

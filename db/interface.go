package db

import "time"

type DBInterface interface {
	SubmitUser(u *User) error
	UpdateUser(address string, role int, stakeTx string, stakeAmount string, expireTime time.Time) error
	GetUser(address string) (*User, error)

	NewRecommendCoin(rc *RecommendCoin) error
	NewRecommendCoinAndCoinInfo(rc *RecommendCoin) error
	GetCoinList(currentPage int, pageSize int) ([]*RecommendCoin, int64, error)

	NewCoinInfo(ci *CoinInfo) error
	GetCoinInfo(uuid string) (*CoinInfo, error)

	GetCoinWithCoinInfo(address string) (*RecommendCoin, error)

	NewTxHistory(tx *TxHistory) error
	GetTxHistoryByAddress(address string) ([]*TxHistory, error)
}

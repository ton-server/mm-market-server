package db

import "time"

type DBInterface interface {
	SubmitUser(u *User) error
	UpdateUser(address string, role int, stakeTx string, stakeAmount string, expireTime time.Time, utime int64) error
	GetUser(address string) (*User, error)
	GetNormalUser() ([]*User, error)

	NewRecommendCoin(rc *RecommendCoin) error
	NewRecommendCoinAndCoinInfo(rc *RecommendCoin) error
	GetCoinList(currentPage int, pageSize int, full bool) ([]*RecommendCoin, int64, error)

	NewCoinInfo(ci *CoinInfo) error
	GetCoinInfo(uuid string) (*CoinInfo, error)

	GetCoinWithCoinInfo(address string) (*RecommendCoin, error)

	NewTxHistory(tx *TxHistory) error
	GetTxHistoryByAddress(address string) ([]*TxHistory, error)

	NewTask(task *Task) error
	GetActiveTask() ([]*Task, error)
	UpdateTask(address string, active int) error

	NewCoinPrice(r *CoinPriceRecord) error
	GetCoinPriceList(address string) (current, pre *CoinPriceRecord, err error)
}

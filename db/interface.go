package db

type DBInterface interface {
	SubmitUser(u *User) error
	UpdateUser(address string, m map[string]any) error
	GetUser(address string) (*User, error)

	NewRecommendCoin(rc *RecommendCoin) error
	GetCoinList(currentPage int, pageSize int) ([]*RecommendCoin, int64, error)

	NewCoinInfo(ci *CoinInfo) error
	GetCoinInfo(uuid string) (*CoinInfo, error)

	GetCoinWithCoinInfo(address string) (*RecommendCoin, error)

	NewTxHistory(tx *TxHistory) error
	GetTxHistoryByAddress(address string) ([]*TxHistory, error)
}

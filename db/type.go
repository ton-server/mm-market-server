package db

import "time"

type User struct {
	Id          int64     `json:"id" gorm:"primary_key;auto_increment"`
	Name        string    `json:"nickname" gorm:"column:nickname"`
	Address     string    `json:"address" gorm:"column:address;unique"`
	Role        int       `json:"role" gorm:"column:role;default:0" ` //1:vip,0:normal,2:ing
	StakeTx     string    `json:"stakeTx" gorm:"column:stake_tx"`
	StakeAmount string    `json:"stakeAmount" gorm:"column:stake_amount"`
	ExpireTime  time.Time `json:"expireTime" gorm:"column:expire_time"`
	ExpireTime2 int64     `json:"expireTime2" gorm:"-"`
	CreateTime  time.Time `json:"createTime" gorm:"column:create_time"`
	UpdateTime  time.Time `json:"updateTime" gorm:"column:update_time"`
}

func (r *User) TableName() string {
	return "table_user"
}

type RecommendCoin struct {
	Id              int64     `json:"id" gorm:"primary_key;auto_increment"`
	UUID            string    `json:"uuid"  gorm:"column:uuid;unique"`
	Name            string    `json:"nickName" gorm:"column:nick_name"`
	Symbol          string    `json:"symbol" gorm:"column:symbol"`
	Decimals        uint8     `json:"decimals" gorm:"column:decimals"`
	TotalSupply     string    `json:"totalSupply" gorm:"column:total_supply"`
	ContractAddress string    `json:"contractAddress" gorm:"column:contract_address"`
	Index           int       `json:"index" gorm:"column:index"`
	CoinInfo        *CoinInfo `json:"coinInfo" gorm:"-"`
	ExpireTime2     int64     `json:"expireTime2" gorm:"-"`
	ExpireTime      time.Time `json:"expireTime" gorm:"column:expire_time"`
	CreateTime      time.Time `json:"createTime" gorm:"column:create_time"`
	UpdateTime      time.Time `json:"updateTime" gorm:"column:update_time"`
}

func (r *RecommendCoin) TableName() string {
	return "table_recommend_coin"
}

type CoinInfo struct {
	Id     int64  `json:"id" gorm:"primary_key;auto_increment"`
	UUID   string `json:"uuid"  gorm:"column:uuid;unique"`
	Detail string `json:"detail" gorm:"column:detail"`
}

func (r *CoinInfo) TableName() string {
	return "table_coin_info"
}

type TxHistory struct {
	Id              int64     `json:"id" gorm:"primary_key;auto_increment"`
	FromAddress     string    `json:"fromAddress" gorm:"column:from_address"`
	ToAddress       string    `json:"toAddress" gorm:"column:to_address"`
	ContractAddress string    `json:"contractAddress" gorm:"column:contract_address"`
	Amount          string    `json:"amount" gorm:"column:amount"`
	TxId            string    `json:"txId" gorm:"column:tx_id"`
	TxStatus        uint8     `json:"txStatus" gorm:"column:tx_status"` //1:交易成功，0:交易失败，2:交易进行中
	TxInfo          string    `json:"txInfo" gorm:"column:tx_info"`
	CreateTime2     int64     `json:"createTime2" gorm:"-"`
	CreateTime      time.Time `json:"createTime" gorm:"column:create_time"`
	UpdateTime      time.Time `json:"updateTime" gorm:"column:update_time"`
}

func (r *TxHistory) TableName() string {
	return "table_tx_history"
}

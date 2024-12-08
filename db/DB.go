package db

import (
	"errors"
	"time"

	"github.com/sunjiangjun/xlog"
	"gorm.io/gorm"
)

const TimeFormat = "2006-01-02 15"

type DB struct {
	core *gorm.DB
	log  *xlog.XLog
}

func (db *DB) NewTask(task *Task) error {
	return db.core.Omit("id", "create_time", "update_time").Create(task).Error
}

func (db *DB) GetActiveTask() ([]*Task, error) {
	var list []*Task
	err := db.core.Model(Task{}).Where("active=? and expire_time>=?", 1, time.Now().UTC()).Scan(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (db *DB) UpdateTask(address string, active int) error {
	m := make(map[string]any, 1)
	m["active"] = active
	return db.core.Model(Task{}).Where("contract_address=?", address).Updates(m).Error
}

func (db *DB) NewCoinPrice(r *CoinPriceRecord) error {
	return db.core.Omit("id", "create_time").Create(r).Error
}

func (db *DB) GetCoinPriceList(address string) (*CoinPriceRecord, *CoinPriceRecord, error) {
	//curr := time.Now().UTC().Format(TimeFormat)
	pre := time.Now().Add(-24 * time.Hour).UTC().Format(TimeFormat)

	var currentRecord CoinPriceRecord

	err := db.core.Model(CoinPriceRecord{}).Where("contract_address=?", address).Order("create_time desc").First(&currentRecord).Error
	if err != nil {
		return nil, nil, err
	}

	var preRecord CoinPriceRecord
	err = db.core.Model(CoinPriceRecord{}).Where("contract_address=? and record_time=?", address, pre).Order("create_time desc").First(&preRecord).Error
	if err != nil {
		return &currentRecord, nil, nil
	}
	return &currentRecord, &preRecord, nil
}

func (db *DB) SubmitUser(u *User) error {
	u.ExpireTime = time.Now().UTC().Add(24 * time.Hour)
	u.Name = ""
	return db.core.Omit("id", "create_time", "update_time").Create(u).Error
}

func (db *DB) UpdateUser(address string, role int, stakeTx string, stakeAmount string, expireTime time.Time, utime int64) error {
	m := make(map[string]any, 4)
	m["role"] = role
	m["stake_tx"] = stakeTx
	m["stake_amount"] = stakeAmount
	m["expire_time"] = expireTime
	m["utime"] = utime
	return db.core.Model(User{}).Where("address=? and utime<?", address, utime).Updates(m).Error
}

func (db *DB) GetNormalUser() ([]*User, error) {
	var list []*User
	err := db.core.Model(User{}).Where("role=? or expire_time<?", 0, time.Now().UTC()).Scan(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (db *DB) GetUser(address string) (*User, error) {
	var u User
	err := db.core.Model(User{}).Where("address=?", address).First(&u).Error
	if err != nil {
		if err.Error() == "record not found" {
			_ = db.SubmitUser(&User{Address: address})
			err = db.core.Model(User{}).Where("address=?", address).First(&u).Error
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	return &u, nil
}

func (db *DB) NewRecommendCoin(rc *RecommendCoin) error {
	return db.core.Omit("id", "create_time", "update_time").Create(rc).Error
}

func (db *DB) NewRecommendCoinAndCoinInfo(rc *RecommendCoin) error {
	err := db.NewCoinInfo(rc.CoinInfo)
	if err != nil {
		return err
	}

	err = db.NewTask(&Task{
		UUID:            rc.UUID,
		ContractAddress: rc.ContractAddress,
		ExpireTime:      rc.ExpireTime,
	})
	if err != nil {
		return err
	}

	err = db.NewRecommendCoin(rc)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) GetCoinList(currentPage int, pageSize int, full bool) ([]*RecommendCoin, int64, error) {
	if currentPage < 1 {
		return nil, 0, errors.New("currentPage more then 1 always")
	}

	var total int64
	err := db.core.Model(RecommendCoin{}).Where("expire_time>=?", time.Now().UTC()).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	var list []*RecommendCoin
	err = db.core.Model(RecommendCoin{}).Where("expire_time>=?", time.Now().UTC()).Order("`index` desc").Offset((currentPage - 1) * pageSize).Limit(pageSize).Scan(&list).Error
	if err != nil {
		return nil, 0, err
	}

	if full {
		for _, v := range list {
			c, err := db.GetCoinInfo(v.UUID)
			if err != nil {
				continue
			}
			v.CoinInfo = c
		}
	}

	return list, total, nil
}

func (db *DB) NewCoinInfo(ci *CoinInfo) error {
	return db.core.Omit("id", "create_time", "update_time").Create(ci).Error
}

func (db *DB) GetCoinInfo(uuid string) (*CoinInfo, error) {
	var c CoinInfo
	err := db.core.Model(CoinInfo{}).Where("uuid=?", uuid).First(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (db *DB) GetCoinWithCoinInfo(uuid string) (*RecommendCoin, error) {
	var c RecommendCoin
	err := db.core.Model(RecommendCoin{}).Where("uuid=?", uuid).First(&c).Error
	if err != nil {
		return nil, err
	}

	coinInfo, err := db.GetCoinInfo(uuid)
	if err != nil {
		return nil, err
	}
	c.CoinInfo = coinInfo
	return &c, nil
}

func (db *DB) NewTxHistory(tx *TxHistory) error {
	return db.core.Omit("id", "create_time", "update_time").Create(tx).Error
}

func (db *DB) GetTxHistoryByAddress(address string) ([]*TxHistory, error) {
	var list []*TxHistory
	err := db.core.Model(TxHistory{}).Where("from_address=? or to_address=?", address, address).Order("create_time desc").Limit(100).Scan(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func NewDB(db *gorm.DB, log *xlog.XLog) *DB {
	return &DB{
		core: db,
		log:  log,
	}
}

package db

import (
	"reflect"
	"testing"

	"github.com/sunjiangjun/xlog"
	"gorm.io/gorm"
)

func TestDB_GetCoinInfo(t *testing.T) {
	type fields struct {
		core *gorm.DB
		log  *xlog.XLog
	}
	type args struct {
		uuid string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *CoinInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			D := &DB{
				core: tt.fields.core,
				log:  tt.fields.log,
			}
			got, err := D.GetCoinInfo(tt.args.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCoinInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCoinInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_GetCoinList(t *testing.T) {
	type fields struct {
		core *gorm.DB
		log  *xlog.XLog
	}
	type args struct {
		currentPage int
		pageSize    int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*RecommendCoin
		want1   int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			D := &DB{
				core: tt.fields.core,
				log:  tt.fields.log,
			}
			got, got1, err := D.GetCoinList(tt.args.currentPage, tt.args.pageSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCoinList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCoinList() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetCoinList() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDB_GetCoinWithCoinInfo(t *testing.T) {
	type fields struct {
		core *gorm.DB
		log  *xlog.XLog
	}
	type args struct {
		address string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *RecommendCoin
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			D := &DB{
				core: tt.fields.core,
				log:  tt.fields.log,
			}
			got, err := D.GetCoinWithCoinInfo(tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCoinWithCoinInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCoinWithCoinInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_GetTxHistoryByAddress(t *testing.T) {
	type fields struct {
		core *gorm.DB
		log  *xlog.XLog
	}
	type args struct {
		address string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *TxHistory
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			D := &DB{
				core: tt.fields.core,
				log:  tt.fields.log,
			}
			got, err := D.GetTxHistoryByAddress(tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTxHistoryByAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTxHistoryByAddress() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_GetUser(t *testing.T) {
	type fields struct {
		core *gorm.DB
		log  *xlog.XLog
	}
	type args struct {
		address string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			D := &DB{
				core: tt.fields.core,
				log:  tt.fields.log,
			}
			got, err := D.GetUser(tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_NewCoinInfo(t *testing.T) {
	type fields struct {
		core *gorm.DB
		log  *xlog.XLog
	}
	type args struct {
		ci *CoinInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			D := &DB{
				core: tt.fields.core,
				log:  tt.fields.log,
			}
			if err := D.NewCoinInfo(tt.args.ci); (err != nil) != tt.wantErr {
				t.Errorf("NewCoinInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDB_NewRecommendCoin(t *testing.T) {
	type fields struct {
		core *gorm.DB
		log  *xlog.XLog
	}
	type args struct {
		rc *RecommendCoin
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			D := &DB{
				core: tt.fields.core,
				log:  tt.fields.log,
			}
			if err := D.NewRecommendCoin(tt.args.rc); (err != nil) != tt.wantErr {
				t.Errorf("NewRecommendCoin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDB_NewTxHistory(t *testing.T) {
	type fields struct {
		core *gorm.DB
		log  *xlog.XLog
	}
	type args struct {
		tx *TxHistory
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			D := &DB{
				core: tt.fields.core,
				log:  tt.fields.log,
			}
			if err := D.NewTxHistory(tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("NewTxHistory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDB_SubmitUser(t *testing.T) {
	type fields struct {
		core *gorm.DB
		log  *xlog.XLog
	}
	type args struct {
		u *User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			D := &DB{
				core: tt.fields.core,
				log:  tt.fields.log,
			}
			if err := D.SubmitUser(tt.args.u); (err != nil) != tt.wantErr {
				t.Errorf("SubmitUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDB_UpdateUser(t *testing.T) {
	type fields struct {
		core *gorm.DB
		log  *xlog.XLog
	}
	type args struct {
		address string
		m       map[string]any
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			D := &DB{
				core: tt.fields.core,
				log:  tt.fields.log,
			}
			if err := D.UpdateUser(tt.args.address, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewDB(t *testing.T) {
	type args struct {
		db  *gorm.DB
		log *xlog.XLog
	}
	tests := []struct {
		name string
		args args
		want *DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDB(tt.args.db, tt.args.log); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

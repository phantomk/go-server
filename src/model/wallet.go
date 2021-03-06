package model

import (
	"time"
)

var (
	WalletCNY  = "CNY"
	WalletUSD  = "USD"
	WalletCOIN = "COIN"
	Wallets    = []string{WalletCNY, WalletUSD, WalletCOIN}
)

type Wallet struct {
	Id        string  `gorm:"primary_key;unique;notnull;index;type:varchar(32)" json:"id"` // 用户ID
	Currency  string  `gorm:"not null;type:varchar(12)" json:"currency"`
	Balance   float64 `gorm:"not null;" json:"balance"` // 可用余额
	Frozen    float64 `gorm:"not null;" json:"frozen"`  // 冻结余额
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// 人民币
type WalletCny struct {
	Wallet
}

// 美元
type WalletUsd struct {
	Wallet
}

// 我们平台自己的币
type WalletCoin struct {
	Wallet
}

func (news *WalletCny) TableName() string {
	return "wallet_cny"
}

func (news *WalletUsd) TableName() string {
	return "wallet_usd"
}

func (news *WalletCoin) TableName() string {
	return "wallet_coin"
}

package models

import (
	"ethereum/config"
	"time"
)

const (
	// Failure 交易失败
	Failure = 0
	// Success 交易成功
	Success = 1
	// Pending 交易确认中
	Pending = 2 // 交易确认中
)

// TransactionLog 交易记录
type TransactionLog struct {
	ID uint `gorm:"primary_key" json:"id"`
	// 交易hash
	TransactionHash string `json:"transaction_hash"`
	// 交易发起人
	From string `json:"from"`
	// 交易接收人
	To string `json:"to"`
	// 发送金额
	Value int64 `json:"value"`
	// 交易类型 ETH/USDT
	Type string `json:"type"`
	// 交易状态
	Status    uint8      `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

// Create 创建交易记录
func (tl *TransactionLog) Create() {
	db := config.GetMysql()
	db.Create(tl)
	db.Close()
}

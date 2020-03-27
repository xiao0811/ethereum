package models

import (
	"ethereum/config"
	"math/big"
	"time"
)

// MonitorLog 交易监控日志结构
type MonitorLog struct {
	ID              uint       `gorm:"primary_key" json:"id"`
	BlockNumber     uint64     `json:"block_number"`
	BlockHash       string     `json:"block_hash"`
	TransactionHash string     `json:"transaction_hash"`
	From            string     `json:"from"`
	To              string     `json:"to"`
	Value           *big.Int   `json:"value"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `sql:"index" json:"deleted_at"`
}

// Create 创建监听日志
func (ml *MonitorLog) Create() {
	db := config.GetMysql()
	defer db.Close()
	db.Create(&ml)
}

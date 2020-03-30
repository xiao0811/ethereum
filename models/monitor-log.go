package models

import (
	"ethereum/config"
	"time"
)

// MonitorLog 交易监控日志结构
type MonitorLog struct {
	ID              uint       `gorm:"primary_key" json:"id"`
	BlockNumber     uint64     `json:"block_number"`
	BlockHash       string     `json:"block_hash" gorm:"type:char(66)"`
	TransactionHash string     `json:"transaction_hash" gorm:"type:char(66)"`
	From            string     `json:"from" gorm:"type:char(42)"`
	To              string     `json:"to" gorm:"type:char(42)"`
	Value           int64      `json:"value"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `sql:"index" json:"deleted_at,omitempty"`
}

// Create 创建监听日志
func (ml *MonitorLog) Create() {
	db := config.GetMysql()
	defer db.Close()
	db.Create(&ml)
}

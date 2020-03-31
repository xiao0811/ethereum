package models

import (
	"ethereum/config"
	"time"
)

// Recharge 充值
type Recharge struct {
	ID    uint   `gorm:"primary_key" json:"id"`
	From  string `json:"from" gorm:"type:char(42)"`
	To    string `json:"to" gorm:"type:char(42)"`
	Value int64  `json:"value"`
	// 状态 models.
	Status uint8 `json:"status" gorm:"default:2"`
	// 回调成功时间
	Callback  time.Time  `json:"callback" gorm:"default:NULL"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
}

// Create 新建充值记录
func (r *Recharge) Create() {
	db := config.GetMysql()
	defer db.Close()
	db.Create(r)
}

// 交易 transaction/recharge

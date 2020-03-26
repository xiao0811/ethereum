package config

import (
	"log"

	// mysql 必要组件
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// GetMysql 获取数据库连接
func GetMysql() *gorm.DB {
	db, err := gorm.Open("mysql", "root:123456@/ethereum?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err)
	}

	return db
}

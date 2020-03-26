package config

import (
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	once sync.Once
)

func GetMysql() *gorm.DB {
	var db *gorm.DB
	var err error
	once.Do(func() {
		db, err = gorm.Open("mysql", "root:123456@/ethereum?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			log.Fatalln(err)
		}
	})
	return db
}

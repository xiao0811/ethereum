package config

import (
	"log"
	"strconv"

	// mysql 必要组件
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// GetMysql 获取数据库连接
func GetMysql() *gorm.DB {
	conf := GetConfig().MysqlConfig
	db, err := gorm.Open("mysql", conf.Username+":"+conf.Password+"@tcp("+
		conf.Host+":"+strconv.Itoa(conf.Port)+")/"+conf.Database+
		"?charset=utf8&parseTime=True&loc=Local",
	)
	if err != nil {
		log.Fatalln(err)
	}

	return db
}

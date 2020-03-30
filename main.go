package main

import (
	"ethereum/config"
	"ethereum/models"
	"ethereum/router"
	"log"
)

func init() {
	db := config.GetMysql()
	// 自动迁移
	err := db.AutoMigrate(
		&models.Addr{},           // 地址库
		&models.TransactionLog{}, // 交易记录
		&models.MonitorLog{},     // 监听记录
	).Error
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
}

func main() {
	app := router.GetRouter()
	if err := app.Run(":8080"); err != nil {
		log.Fatalln("服务器启动失败:", err)
	}
}

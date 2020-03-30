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
	db.AutoMigrate(&models.Addr{}, &models.TransactionLog{}, &models.MonitorLog{})
}

func main() {
	app := router.GetRouter()
	if err := app.Run(":8080"); err != nil {
		log.Fatalln("服务器启动失败:", err)
	}
}

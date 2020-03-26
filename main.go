package main

import (
	"ethereum/router"
	"log"
)

func main() {
	app := router.GetRouter()
	if err := app.Run(":8080"); err != nil {
		log.Fatalln("服务器启动失败:", err)
	}
}

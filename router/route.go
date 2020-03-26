package router

import (
	"ethereum/controllers/account"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	app := gin.Default()

	// 根据地址查看用户账户
	app.POST("balance", account.Balance)
	return app
}

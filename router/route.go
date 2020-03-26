package router

import (
	"ethereum/controllers/account"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	app := gin.Default()

	// 根据地址查看用户账户
	app.POST("balance", account.Balance)
	// 代币转账
	app.POST("tokenTransfer", account.TokenTransfer)
	// 生成账户
	app.POST("accountGenerate", account.Generate)
	// ETH 转账
	app.POST("ETHTransfer", account.ETHTransfer)
	return app
}

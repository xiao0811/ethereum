package router

import (
	"ethereum/controllers/account"
	"ethereum/middleware"

	"github.com/gin-gonic/gin"
)

// GetRouter 返回所有路由
func GetRouter() *gin.Engine {
	app := gin.Default()

	v1 := app.Group("v1")
	{
		// 使用RSA中间件
		v1.Use(middleware.Rsa())
		// 根据地址查看用户账户
		v1.POST("balance", account.Balance)
		// 代币转账
		v1.POST("tokenTransfer", account.TokenTransfer)
		// 生成账户
		v1.POST("accountGenerate", account.Generate)
		// ETH 转账
		v1.POST("ETHTransfer", account.ETHTransfer)
		// 获取一个用户
		v1.POST("createAccount", account.NewAccount)
	}

	return app
}

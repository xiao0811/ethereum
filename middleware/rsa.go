package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Rsa 中间件
func Rsa() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("xiaosha")
		c.Next()
	}
}

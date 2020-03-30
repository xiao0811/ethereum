package middleware

import (
	"github.com/gin-gonic/gin"
)

// Rsa 中间件
func Rsa() gin.HandlerFunc {
	return func(c *gin.Context) {
		// _data := c.PostForm("data")
		// // 密文
		// ciphertext, err := base64.StdEncoding.DecodeString(_data)
		// if err != nil {
		// 	handles.Error(http.StatusBadRequest, "数据加密出错", c)
		// }
		// data, err := config.RsaDecrypt(ciphertext)
		// if err != nil {
		// 	handles.Error(http.StatusBadRequest, "数据加密出错", c)
		// }
		// c.Set("data", data)
		c.Next()
	}
}

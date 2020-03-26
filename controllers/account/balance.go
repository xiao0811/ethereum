package account

import (
	"ethereum/handles"
	"ethereum/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Balance 获取用户账户信息
func Balance(c *gin.Context) {
	addr := c.PostForm("address")
	// 验证地址
	if !handles.AddressVerify(addr) {
		handles.Error(http.StatusBadRequest, "用户地址不正确", c)
	} else {
		handles.Success("OK", models.Addr{Addr: addr}.GetBalance(), c)
	}
}

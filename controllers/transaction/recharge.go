package transaction

import (
	"encoding/json"
	"ethereum/handles"
	"ethereum/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRecharge(c *gin.Context) {
	var recharge models.Recharge
	data, _ := c.Get("data")
	if err := json.Unmarshal(data.([]byte), &recharge); err != nil {
		handles.Error(http.StatusBadRequest, "数据格式不正确", c)
		return
	}
	recharge.Create()
	// 从总账户给用户提供外部账户转账
	handles.Success("提交成功!", gin.H{
		"data": "0x58a788d69e0f80c3f6bbd28b687bf93c259c594eabaf54ff5b1ca7479d4fea47",
	}, c)
}

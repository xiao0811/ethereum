package handles

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ReturnJson 返回json
func (rd ResponseData) ReturnJson(c *gin.Context) {
	c.AbortWithStatusJSON(rd.Code, rd)
}

// Error 错误返回
func Error(code int, message string, c *gin.Context) {
	ResponseData{
		Code:    code,
		Message: message,
		Data:    nil,
	}.ReturnJson(c)
}

// Success 正确返回
func Success(message string, data interface{}, c *gin.Context) {
	ResponseData{
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	}.ReturnJson(c)
}

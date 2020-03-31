package account

import (
	"ethereum/config"
	"ethereum/handles"
	"ethereum/models"
	"log"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/gin-gonic/gin"
)

// Generate 批量生成用户地址
func Generate(c *gin.Context) {
	_number := c.DefaultPostForm("number", "5")
	number, err := strconv.Atoi(_number)
	if err != nil {
		handles.Error(http.StatusBadRequest, "个数不正确", c)
	} else {
		for i := 0; i < number; i++ {
			ks := keystore.NewKeyStore(
				"./tmp",
				keystore.StandardScryptN,
				keystore.StandardScryptP,
			)
			account, err := ks.NewAccount(config.GetConfig().AddrConfig.Password)
			if err != nil {
				log.Fatal(err)
			}
			db := config.GetMysql()
			addr := models.Addr{
				Addr: account.Address.String(),
				URL:  account.URL.String(),
			}
			db.Create(&addr)
			db.Close()
		}
	}
}

// NewAccount 从地址从拿出一个地址
func NewAccount(c *gin.Context) {
	var addr models.Addr
	db := config.GetMysql()
	if err := db.Where("status = ?", 0).First(&addr).Error; err != nil {
		handles.Error(http.StatusInternalServerError, "服务器没地址了", c)
		return
	}
	addr.Status = 1
	db.Save(&addr)
	db.Close()
	handles.Success("OK", gin.H{"addr": addr.Addr}, c)
}

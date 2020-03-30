package models

import (
	"context"
	"ethereum/config"
	"ethereum/contract/wrapper/eztoken"
	"io/ioutil"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

// Addr 地址库
// status默认值为0, 表示地址暂未分配用, 1 表示已分配
type Addr struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Addr      string     `json:"addr" gorm:"unique;not null;type:char(42)"`
	URL       string     `json:"url"`
	Status    uint8      `json:"status" gorm:"default:0"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
}

// GetBalance 获取用户账户信息
func (addr Addr) GetBalance() gin.H {
	client, err := ethclient.Dial("http://47.244.209.218:8545")
	if err != nil {
		log.Fatal(err)
	}

	// 合约地址
	// Golem (GNT) Address less  0xe78A0F7E598Cc8b0Bb87894B0F60dD2a88d6a8Ab
	tokenAddress := common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7")
	instance, err := eztoken.NewToken(tokenAddress, client)
	if err != nil {
		log.Fatalln(err)
	}

	// 查询账户地址
	address := common.HexToAddress(addr.Addr)

	// ETH
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Token
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}

	return gin.H{
		"eth":   balance,
		"token": bal.Int64(),
	}
}

// GetAccountPrivateKey 获取用户PrivateKey
func (addr *Addr) GetAccountPrivateKey() string {
	db := config.GetMysql()
	db.Where("addr = ?", addr.Addr).First(addr)
	defer db.Close()
	return importKs(addr.URL[11:])
}

func importKs(file string) string {
	keyJSON, err := ioutil.ReadFile(file)
	if err != nil {
		return ""
	}
	key, _ := keystore.DecryptKey(
		keyJSON,                                // keystore json
		config.GetConfig().AddrConfig.Password, // 解密口令，对称
	)
	return hexutil.Encode(crypto.FromECDSA(key.PrivateKey))
}

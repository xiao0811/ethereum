package models

import (
	"context"
	"ethereum/config"
	"ethereum/contract/wrapper/eztoken"
	"log"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Addr struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Addr      string     `json:"addr" gorm:"unique;not null;size:42"`
	URL       string     `json:"url"`
	Status    uint8      `json:"status" gorm:"default:0"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

func (addr Addr) GetBalance() gin.H {
	client, err := ethclient.Dial(config.SERVER)
	if err != nil {
		log.Fatal(err)
	}

	// 合约地址
	// Golem (GNT) Address less  0xe78A0F7E598Cc8b0Bb87894B0F60dD2a88d6a8Ab
	tokenAddress := common.HexToAddress(config.TokenAddress)
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

	// name, err := instance.Name(&bind.CallOpts{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// symbol, err := instance.Symbol(&bind.CallOpts{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// decimals, err := instance.Decimals(&bind.CallOpts{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// fmt.Printf("name: %s\n", name)         // "name: Golem Network"
	// fmt.Printf("symbol: %s\n", symbol)     // "symbol: GNT"
	// fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"
	return gin.H{
		"eth":   balance,
		"token": bal.Int64(),
	}
}

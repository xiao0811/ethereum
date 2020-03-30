package config

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

// GetClient 获取以太坊客户端连接
func GetClient() *ethclient.Client {
	client, err := ethclient.Dial(GetConfig().InfuraConfig.Http)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/9831d331dbcd429b88b990811cffd50e")
	if err != nil {
		log.Fatal(err)
	}

	// 合约地址
	contractAddress := common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7")
	query := ethereum.FilterQuery{
		// FromBlock: big.NewInt(2394201),
		// ToBlock:   big.NewInt(2394201),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			tx, isPending, err := client.TransactionByHash(context.Background(), vLog.TxHash)
			if err != nil {
				log.Fatal(err)
			}

			if !isPending {
				fmt.Println("txhash:", tx.Hash().Hex())    // txhash
				fmt.Println("To:", tx.To().String())       // to
				fmt.Println("Value:", tx.Value().String()) // Value
			}
		}
	}
}

package main

import (
	"ethereum/config"
	"ethereum/contract/wrapper/eztoken"
	"ethereum/ethtool"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	log.Println("监听交易开始......")
	monitor()
}

func monitor() {
	// 合约地址
	contractAddress := common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7")

	client, err := ethtool.Dial(config.WSSSERVER)
	assert(err)

	inst, err := eztoken.NewEztoken(contractAddress, client)
	assert(err)

	watchOpts := &bind.WatchOpts{}
	events := make(chan *eztoken.EztokenTransfer)
	var _from []common.Address
	var _to []common.Address
	sub, err := inst.WatchTransfer(watchOpts, events, _from, _to)
	assert(err)

	for {
		select {
		case err := <-sub.Err():
			log.Println(err)
		case event := <-events:
			fmt.Println("captured:")
			fmt.Println("-> from: ", event.From.Hex())
			fmt.Println("-> to: ", event.To.Hex())
			fmt.Println("-> value:", event.Value)
			fmt.Println("-> TxHash:", event.Raw.TxHash.String())
			// 处理交易逻辑
		}
	}
}

func assert(err error) {
	if err != nil {
		log.Println(err)
	}
}

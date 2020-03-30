package main

import (
	"ethereum/config"
	"ethereum/contract/wrapper/eztoken"
	"ethereum/ethtool"
	"ethereum/models"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func init() {
	db := config.GetMysql()
	db.AutoMigrate(&models.MonitorLog{})
	db.Close()
}

func main() {
	log.Println("监听交易开始......")
	monitor()
}

func monitor() {
	// 合约地址
	conf := config.GetConfig()
	contractAddress := common.HexToAddress(conf.UsdtConfig.Token)

	client, err := ethtool.Dial(conf.InfuraConfig.Wss)
	assert(err)
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
			// fmt.Println(event.Raw.BlockHash.String(), ":")
			// fmt.Println("from: ", event.From.Hex())
			// fmt.Println("to: ", event.To.Hex())
			// fmt.Println("value:", event.Value)
			// fmt.Println("TxHash:", event.Raw.TxHash.String())
			// fmt.Println("BlockNumber:", event.Raw.BlockNumber)
			// fmt.Println("---------------------------------------------------")
			// 处理交易逻辑
			monitorLog := &models.MonitorLog{
				BlockNumber:     event.Raw.BlockNumber,
				BlockHash:       event.Raw.BlockHash.String(),
				TransactionHash: event.Raw.TxHash.String(),
				From:            event.From.Hex(),
				To:              event.To.Hex(),
				Value:           event.Value.Int64(),
			}
			monitorLog.Create()
		}
	}
}

func assert(err error) {
	if err != nil {
		log.Println(err)
	}
}

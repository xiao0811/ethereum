package main

import (
	"ethereum/config"
	"ethereum/contract/wrapper/eztoken"
	"ethereum/ethtool"
	"ethereum/models"
	"fmt"
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
			fmt.Println(event.Raw.BlockHash.String(), ":")
			fmt.Println("from: ", event.From.Hex())
			fmt.Println("to: ", event.To.Hex())
			fmt.Println("value:", event.Value.Int64())
			fmt.Println("TxHash:", event.Raw.TxHash.String())
			fmt.Println("BlockNumber:", event.Raw.BlockNumber)
			fmt.Println("---------------------------------------------------")
			// 处理交易逻辑
			// 是否是充值
			// go recharge(event.To.Hex(), event.Raw.TxHash.String(), event.Value.Int64())
			// 是否提币
			// go withdrawal(event.To.Hex(), event.From.Hex(), event.Raw.TxHash.String(), event.Value.Int64())
		}
	}
}

func assert(err error) {
	if err != nil {
		log.Println(err)
	}
}

// recharge 充值
// 内部地址 : addrs库中 status = 1
// 保存交易hash
func recharge(to, TxHash string, value int64) {
	// 1, 判断 to 是否是内部地址
	var internalAddress []models.Addr
	db := config.GetMysql()
	db.Where("status = ?", 1).Find(&internalAddress)
	for _, addr := range internalAddress {
		if to == addr.Addr {
			// 2, 内部地址 写入充值记录
		}
	}
}

// withdrawal 提币
func withdrawal(to, from, TxHash string, value int64) {
	// 1, from 是否为指定转币地址
	if from != "" {
		return
	}
	// 2, to 是否为提币库中地址

	// 3, 验证value 写入库中
}

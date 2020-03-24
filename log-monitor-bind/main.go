package main

import (
	"ethereum/config"
	"ethereum/contract/wrapper/eztoken"
	"ethereum/ethtool"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	addrHexBytes, err := ioutil.ReadFile("../contract/build/EzToken.addr")
	assert(err)
	contractAddress := common.HexToAddress(string(addrHexBytes))
	assert(err)
	client, err := ethtool.Dial("ws://" + config.SERVER)
	assert(err)

	inst, err := eztoken.NewEztoken(contractAddress, client)
	assert(err)

	watchOpts := &bind.WatchOpts{}
	events := make(chan *eztoken.EztokenTransfer)
	var _from []common.Address
	var _to []common.Address
	sub, err := inst.WatchTransfer(watchOpts, events, _from, _to)
	assert(err)
	// fmt.Println(sub)

	for {
		select {
		case err := <-sub.Err():
			panic(err)
		case event := <-events:
			fmt.Println("captured:")
			fmt.Println("-> from: ", event.From.Hex())
			fmt.Println("-> to: ", event.To.Hex())
			fmt.Println("-> value:", event.Value)
		}
	}
}

func assert(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

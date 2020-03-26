package main

import (
	"ethereum/models"
	"ethereum/router"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const Password = "ethereum"

func main() {
	// db := config.GetMysql()
	// var addrs []models.Addr
	//
	// db.Find(&addrs)
	//
	// for _, addr := range addrs {
	// 	// fmt.Println(addr.Addr)
	// 	client, err := ethclient.Dial("http://47.244.209.218:8559")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	//
	// 	// Golem (GNT) Address
	// 	tokenAddress := common.HexToAddress("0xe78A0F7E598Cc8b0Bb87894B0F60dD2a88d6a8Ab")
	// 	instance, err := eztoken.NewToken(tokenAddress, client)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	//
	// 	address := common.HexToAddress(addr.Addr)
	// 	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	balance, err := client.BalanceAt(context.Background(), common.HexToAddress(addr.Addr), nil)
	// 	fmt.Println(addr.Addr, bal.Int64(), balance.Int64())
	// 	fmt.Println(addr.URL)
	// }
	// privateKey()
	app := router.GetRouter()
	app.Run(":8080")
}

func privateKey() {
	file := "/Users/panxiao/Desktop/ethereum/tmp/UTC--2020-03-25T03-05-49.909562000Z--e2da77a35e6a4622433f3d9cb479452fb78e5681"
	keyjson, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalln(err)
	}
	key, _ := keystore.DecryptKey(
		keyjson,  // keystore json
		Password, // 解密口令，对称
	)
	fmt.Println("private key: ", hexutil.Encode(crypto.FromECDSA(key.PrivateKey)))
}

func createKs() {
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(Password)
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open("mysql", "root:123456@/ethereum?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err)
	}
	addr := models.Addr{
		Addr:   account.Address.String(),
		URL:    account.URL.String(),
		Status: 0,
	}
	db.Create(&addr)
	db.Close()
}

func importKs(file string) {
	// file := "./tmp/UTC--2018-07-04T09-58-30.122808598Z--20f8d42fb0f667f2e53930fed426f225752453b3"
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	password := "secret"
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3

	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}
}

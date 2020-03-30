package account

import (
	"context"
	"crypto/ecdsa"
	"ethereum/config"
	"ethereum/contract/wrapper/eztoken"
	"ethereum/ethtool"
	"ethereum/handles"
	"ethereum/models"
	"log"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
)

// Transfer 交易结构体
type Transfer struct {
	From  string `json:"from" form:"from" binding:"required"`
	To    string `json:"to" form:"to" binding:"required"`
	Token int64  `json:"token" form:"token" binding:"required"`
}

// Balance 获取用户账户信息
func Balance(c *gin.Context) {
	addr := c.PostForm("address")
	// 验证地址
	if !handles.AddressVerify(addr) {
		handles.Error(http.StatusBadRequest, "用户地址不正确", c)
	} else {
		handles.Success("OK", models.Addr{Addr: addr}.GetBalance(), c)
	}
}

// TokenTransfer 代币转账
func TokenTransfer(c *gin.Context) {
	var transfer Transfer
	if err := c.ShouldBind(&transfer); err != nil {
		handles.Error(http.StatusBadRequest, "数据有误!", c)
	} else {
		client, err := ethtool.Dial(config.GetConfig().InfuraConfig.Http)
		handles.ErrorVerify(err)
		// 合约地址
		contractAddress := common.HexToAddress(config.GetConfig().UsdtConfig.Token)
		inst, err := eztoken.NewEztoken(contractAddress, client)
		handles.ErrorVerify(err)

		var addr = &models.Addr{Addr: transfer.From}
		privateKey := addr.GetAccountPrivateKey()
		balance := addr.GetBalance()
		var eth = balance["eth"].(*big.Int)
		var token = balance["token"].(int64)
		if eth.Int64() <= 800000000000000 {
			handles.Error(http.StatusBadRequest, "ETH gas不足", c)
			return
		}

		if token < transfer.Token {
			handles.Error(http.StatusBadRequest, "余额不足", c)
			return
		}
		if privateKey == "" {
			handles.Error(http.StatusBadRequest, "用户账户有问题", c)
			return
		}
		// from 私钥
		credential, err := ethtool.HexToCredential(privateKey)
		handles.ErrorVerify(err)

		txOpts := bind.NewKeyedTransactor(credential.PrivateKey)

		// to 地址
		toAddress := common.HexToAddress(transfer.To)
		amount := big.NewInt(transfer.Token)

		tx, err := inst.Transfer(txOpts, toAddress, amount)
		if err != nil {
			handles.Error(http.StatusBadRequest, "转账失败!", c)
		} else {
			handles.Success("转账成功!", gin.H{
				"txid": tx.Hash().Hex(),
			}, c)
			transactionLog := &models.TransactionLog{
				To:              transfer.To,
				From:            transfer.From,
				Value:           transfer.Token,
				Type:            "USDT",
				Status:          models.Pending,
				TransactionHash: tx.Hash().Hex(),
			}
			transactionLog.Create()
		}
	}
}

// ETHTransfer ETH 转账
func ETHTransfer(c *gin.Context) {
	var transfer Transfer
	client := config.GetClient()
	if err := c.ShouldBind(&transfer); err != nil {
		handles.Error(http.StatusBadRequest, "数据有误!", c)
		return
	}

	addr := &models.Addr{Addr: transfer.From}
	_privateKey := addr.GetAccountPrivateKey()
	// _privateKey := "0x4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d"
	privateKey, err := crypto.HexToECDSA(_privateKey[2:])
	if err != nil {
		log.Println(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	value := big.NewInt(transfer.Token) // in wei (1 eth)
	gasLimit := uint64(21000)           // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress(transfer.To)
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	transactionLog := &models.TransactionLog{
		To:              transfer.To,
		From:            transfer.From,
		Value:           transfer.Token,
		Type:            "ETH",
		Status:          models.Pending,
		TransactionHash: signedTx.Hash().Hex(),
	}
	transactionLog.Create()
	handles.Success("OK", gin.H{"transaction_hash": signedTx.Hash().Hex()}, c)
}

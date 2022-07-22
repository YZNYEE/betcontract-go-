package utils

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

const host string = "http://localhost:7545"

var Client *ethclient.Client

func Getclient(host string) *ethclient.Client {
	client, err := ethclient.Dial(host)
	if err != nil {
		log.Fatal("连接结点出错:", err)
	}
	return client
}

func GetAddress(pkey string) common.Address {
	privateKey, err := crypto.HexToECDSA(pkey)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	pubkey, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	address := crypto.PubkeyToAddress(*pubkey)
	return address

}

func Getnouce(add common.Address, client *ethclient.Client) *big.Int {
	nouce, err := client.NonceAt(context.Background(), add, nil)
	if err != nil {
		log.Fatal("获得nouce失败:", err)
	}
	return big.NewInt(int64(nouce))
}
func GetGasPrice(client *ethclient.Client) *big.Int {
	price, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("获得price错误:", err)
	}
	return price
}

func init() {
	Client = Getclient(host)
}

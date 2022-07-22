package utils

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func GetAuthfromprivatekey(client *ethclient.Client, pkey string) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA(pkey)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	_, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	chid, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	topt, err := bind.NewKeyedTransactorWithChainID(privateKey, chid)
	if err != nil {
		log.Fatal(err)
	}

	return topt
}

func Getauth(pkey string, client *ethclient.Client) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA(pkey)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	auth := GetAuthfromprivatekey(client, pkey)
	nouce := Getnouce(fromAddress, client)
	auth.Nonce = nouce
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(400000)
	auth.GasPrice = GetGasPrice(client)
	return auth
	//if err != nil {
	//	log.Fatal(err)
	//}
	//input := "1.0"
	//address, tx, instance, er := bind.DeployContract()
}

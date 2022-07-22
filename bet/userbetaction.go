package bet

import (
	"betcontract/utils"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
)

func Deploycontract(privatekey string) {
	auth := utils.Getauth(privatekey, utils.Client)
	add := utils.GetAddress(privatekey)
	auth.Nonce = utils.Getnouce(add, utils.Client)
	auth.GasPrice = utils.GetGasPrice(utils.Client)
	auth.GasLimit = 6720000
	auth.Value = big.NewInt(400000000)

	contractaddress, contracttx, _, err := DeployBet(auth, utils.Client)
	if err != nil {
		log.Fatal("合约发布失败:", err)
	}
	fmt.Println("合约地址:", contractaddress.Hex())
	fmt.Println("交易地址:", contracttx.Hash().Hex())
}

func Getbetcontract(contractaddress string) *Bet {
	add := common.HexToAddress(contractaddress)
	bet, err := NewBet(add, utils.Client)
	if err != nil {
		log.Fatal("获取合约失败:", err)
	}
	return bet
}

func getopts(pkey string, value int64) *bind.TransactOpts {
	auth := utils.Getauth(pkey, utils.Client)
	add := utils.GetAddress(pkey)
	auth.Nonce = utils.Getnouce(add, utils.Client)
	auth.GasPrice = utils.GetGasPrice(utils.Client)
	auth.GasLimit = 400000
	auth.Value = big.NewInt(value)
	return auth
}

func Startbet(bet *Bet, pkey string) {
	opts := utils.Getauth(pkey, utils.Client)
	tx, err := bet.Start(opts)
	if err != nil {
		log.Fatal("开始赌局失败:", err)
	}
	fmt.Println("交易地址:", tx.Hash().Hex())
}

func Inject(bet *Bet, pkey string, value int64) {
	opts := getopts(pkey, value)
	tx, err := bet.Inject(opts)
	if err != nil {
		log.Fatal("注资失败:", err)
	}
	fmt.Println("交易地址:", tx.Hash().Hex())
}

func Openbet(bet *Bet, pkey string) {
	opts := getopts(pkey, 0)
	tx, err := bet.Open(opts)
	if err != nil {
		log.Fatal("开始赌局失败:", err)
	}
	fmt.Println("交易地址:", tx.Hash().Hex())
}
func Changeowner(bet *Bet, pkey string, next common.Address) {
	opts := getopts(pkey, 0)
	bet.Changeowner(opts, next)
}
func Joinbet(bet *Bet, pkey string, value int64, expect int64) {
	opts := getopts(pkey, value)
	tx, err := bet.Joinbet(opts, big.NewInt(expect))
	if err != nil {
		log.Fatal("参加赌局失败:", err)
	}
	fmt.Println("交易地址:", tx.Hash().Hex())
}

func Getbetbalance(bet *Bet, pkey string) *big.Int {
	//opts := getopts(pkey, 0)
	balance, err := bet.Getbalance(nil)
	if err != nil {
		log.Fatal("获取账户余额出错:", err)
	}
	return balance
}

func Getbetstate(bet *Bet, pkey string) bool {
	//opt := utils.Getauth(pkey, utils.Client)
	state, err := bet.Getstate(nil)
	if err != nil {
		log.Fatal("获取状态失败:", err)
	}
	return state
}

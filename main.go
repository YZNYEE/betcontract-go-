package main

import (
	"betcontract/bet"
	"fmt"
)

const pkey = "8e306c90201c773311201f868273e09d9881fb898a1667e52463f27d9eb1396d"
const player1 = "610fa4ba3409d80b41f9489a65ba36a3e8a44fb12620bff87b0da37987ac7982"
const player2 = "74e4d1284585c6cc00a008790f7f8e1a92d75a525de3e658f6d8e9aab99ff194"
const contractaddress = "0x91FeE6f000C38515189784212A0532A97a422031"

func main() {
	//bet.Deploycontract(pkey)
	b := bet.Getbetcontract(contractaddress)
	//bet.Startbet(b, pkey)
	//state := bet.Getbetstate(b, pkey)
	//fmt.Println(state)
	banlance := bet.Getbetbalance(b, pkey)
	fmt.Println(banlance)
	//bet.Joinbet(b, player1, 10000, 1)
	//time.Sleep(40 * time.Second)
	//banlance = bet.Getbetbalance(b, pkey)
	//fmt.Println(banlance)
	//time.Sleep(40 * time.Second)
	bet.Openbet(b, pkey)
	//bet.Joinbet(b, player2, 20000, 2)
	banlance = bet.Getbetbalance(b, pkey)
	fmt.Println(banlance)

}

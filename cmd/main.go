package main

import (
	"fmt"

	. "github.com/martinmsb/UTX0-Blockchain/cmd/blockchain"
)

func main() {
	var blockchain *Blockchain = new(Blockchain)
	blockchain.InitGenesisBlock("myadress0000000")
	fmt.Println(blockchain.ToString())
	fmt.Print(blockchain.Blocks[0].ToString())
	//fmt.Println(transaction.ToString())
}

package main

import (
	"fmt"

	"github.com/martinmsb/UTX0-Blockchain/internal/blockchain"
)

func main() {
	var blockchain *blockchain.Blockchain = new(blockchain.Blockchain)
	blockchain.InitGenesisBlock("adressvalidator")
	fmt.Println(blockchain.ToString())
	fmt.Print(blockchain.Blocks[0].ToString())
	//fmt.Println(transaction.ToString())
}

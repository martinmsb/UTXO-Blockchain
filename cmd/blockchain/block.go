package blockchain

import (
	"github.com/davecgh/go-spew/spew"
)

type Block struct {
	Index        int
	Timestamp    string
	Hash         string
	PrevHash     string
	Transactions []Transaction
}

func (block *Block) ToString() string {
	return spew.Sdump(block)
}

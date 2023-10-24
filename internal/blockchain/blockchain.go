package blockchain

import (
	"time"

	"github.com/davecgh/go-spew/spew"
)

type Blockchain struct {
	Blocks []Block
	Length int
}

func (bchain *Blockchain) CreateBlock() {

}

func (bchain *Blockchain) InitGenesisBlock(validator string) {
	var input *Input = new(Input)
	input.Value = 100500
	inputs := []Input{
		*input,
	}
	var output1 *Output = new(Output)
	output1.Value = 100000
	output1.ScriptPubKey = "adress00000"
	outputs := []Output{
		*output1,
	}
	var transaction *Transaction = new(Transaction)
	transaction.Inputs = inputs
	transaction.Outputs = outputs
	transaction.Locktime = "0"
	transaction.Reward(validator)
	transaction.GenerateIndex()
	transactions := []Transaction{
		*transaction,
	}
	var block *Block = new(Block)
	block.Index = 0
	block.Timestamp = time.Now().String()
	block.Hash = "TODOHASH"
	block.Transactions = transactions
	blocks := []Block{
		*block,
	}
	bchain.Blocks = blocks
	bchain.Length = 1
}

func (blockchain *Blockchain) ToString() string {
	return spew.Sdump(blockchain)
}

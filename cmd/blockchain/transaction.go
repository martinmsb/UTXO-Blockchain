package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/davecgh/go-spew/spew"
)

type Input struct {
	PrevTxIndex string
	Vout        int
	Value       float64
	ScriptSig   string
}

type Output struct {
	Value        float64
	ScriptPubKey string
}

type Transaction struct {
	TxId     string
	Inputs   []Input
	Outputs  []Output
	Locktime string
}

func (transaction *Transaction) GenerateIndex() {
	record := time.Now().String() + spew.Sdump(transaction.Inputs) + spew.Sdump(transaction.Outputs)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	h2 := sha256.New()
	h2.Write([]byte(hashed))
	hashed2 := h2.Sum(nil)
	transaction.TxId = hex.EncodeToString(hashed2)
}

func (transaction *Transaction) NewUTXO(addressSend string, addresRcv string, value float64) {

}

func (transaction *Transaction) Reward(validator string) {
	var outrew *Output = new(Output)
	outrew.Value = 500
	outrew.ScriptPubKey = "validator"
	transaction.Outputs = append(transaction.Outputs, *outrew)
}

func (transaction *Transaction) ToString() string {
	return spew.Sdump(transaction)
}

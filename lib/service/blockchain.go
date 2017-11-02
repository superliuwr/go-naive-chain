package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/superliuwr/go-naive-chain/lib/data"
	"github.com/superliuwr/go-naive-chain/lib/util"
)

// Blockchain defines a blockchain service
type Blockchain struct {
	Blocks              []data.Block
	currentTransactions []data.Transaction
}

// NewBlockchain returns a new instance of BlockChain
func NewBlockchain() *Blockchain {
	chain := Blockchain{
		Blocks:              []data.Block{},
		currentTransactions: []data.Transaction{},
	}

	chain.AddBlock(100, "Genesis-Block-Hash")

	return &chain
}

// AddBlock creates a new block and adds it to the chain
func (b *Blockchain) AddBlock(proof int, previousHash string) (*data.Block, error) {
	chainLength := len(b.Blocks)

	if len(previousHash) == 0 {
		if chainLength == 0 {
			return nil, fmt.Errorf("unable to add new block: genesis block is missing previousHash value")
		}

		bytes, err := json.Marshal(b.Blocks[chainLength-1])
		if err != nil {
			return nil, fmt.Errorf("unable to add new block: %s", err.Error())
		}

		previousHash = util.HashSha256(bytes)
	}

	block := data.Block{
		Index:        len(b.Blocks) + 1,
		PreviousHash: previousHash,
		Proof:        proof,
		Timestamp:    time.Now(),
		Transactions: []data.Transaction{},
	}

	block.Transactions = append(block.Transactions, b.currentTransactions...)
	b.currentTransactions = []data.Transaction{}

	b.Blocks = append(b.Blocks, block)

	return &block, nil
}

// AddTransaction creates a new transaction and adds it to current transaction list
func (b *Blockchain) AddTransaction(tx data.Transaction) (int, error) {
	lastBlock, err := b.LastBlock()
	if err != nil {
		return 0, fmt.Errorf("unable to add transaction: %s", err.Error())
	}

	b.currentTransactions = append(b.currentTransactions, tx)

	return lastBlock.Index + 1, nil
}

// LastBlock returns the last block of the chain
func (b *Blockchain) LastBlock() (*data.Block, error) {
	length := len(b.Blocks)

	if length > 0 {
		return &b.Blocks[length-1], nil
	}

	return nil, fmt.Errorf("there is no blocks in the chain")
}

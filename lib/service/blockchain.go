package service

import (
	"fmt"	
	"time"

	"github.com/superliuwr/go-naive-chain/lib/data"
	"github.com/superliuwr/go-naive-chain/lib/util"
)

// BlockChain defines a blockchain implementation
type BlockChain struct {
	blocks []data.Block
	currentTransactions []data.Transaction
}

// NewBlockChain returns a new instance of BlockChain
func NewBlockChain() *BlockChain {
	chain := BlockChain{
		blocks: []data.Block{},
		currentTransactions: []data.Transaction{},
	}

	chain.AddBlock(100, "Genesis-Block-Hash")

	return &chain
}

// AddBlock creates a new block and adds it to the chain
func (b *BlockChain) AddBlock(proof int, previousHash string) (*data.Block, error) {
	chainLength := len(b.blocks)
	
	if len(previousHash) == 0 {
		if chainLength ==0 {
			return nil, fmt.Errorf("unable to add new block: genesis block is missing previousHash value")
		}

		hashBytes, err := util.Hash(b.blocks[chainLength - 1])
		if err != nil {
			return nil, fmt.Errorf("unable to add new block: %s", err.Error())
		}

		previousHash = string(hashBytes)
	}

	block := data.Block {
		Index: len(b.blocks) + 1,
		PreviousHash: previousHash,
		Proof: proof,
		Timestamp: time.Now(),
		Transactions: []data.Transaction{},
	}

	block.Transactions = append(block.Transactions, b.currentTransactions...)
	b.currentTransactions = []data.Transaction{}

	b.blocks = append(b.blocks, block)

	return &block, nil
}

// AddTransaction creates a new transaction and adds it to current transaction list
func (b *BlockChain) AddTransaction(sender string, recipient string, amount int) (int, error) {
	lastBlock, err := b.LastBlock()
	if err != nil {
		return 0, fmt.Errorf("unable to add transaction: %s", err.Error())
	}

	b.currentTransactions = append(b.currentTransactions, data.Transaction{
		Sender: sender,
		Recipient: recipient,
		Amount: amount,
	})

	return lastBlock.Index + 1, nil
}

// LastBlock returns the last block of the chain
func (b *BlockChain) LastBlock() (*data.Block, error) {
	length := len(b.blocks)

	if length > 0 {
		return &b.blocks[length - 1], nil
	}
	
	return nil, fmt.Errorf("there is no blocks in the chain")
}
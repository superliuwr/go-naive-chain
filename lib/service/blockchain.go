package service

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"github.com/davecgh/go-spew/spew"
)

// Block represents each 'item' in the blockchain
type Block struct {
	Index     int
	Timestamp string
	Payload   int
	Hash      string
	PrevHash  string
}

// BlockchainService defines a blockchain service
type BlockchainService struct {
	chain []*Block
}

// NewBlockchainService returns a new instance of BlockchainService
func NewBlockchainService() *BlockchainService {
	service := BlockchainService{
		chain: newBlockchain(),
	}

	return &service
}

// Add adds a new block for the given payload
func (s *BlockchainService) Add(payload int) (Block, error) {
	newBlock := s.generateBlock(s.chain[len(s.chain)-1], payload)

	if isBlockValid(&newBlock, s.chain[len(s.chain)-1]) {
		newBlockchain := append(s.chain, &newBlock)
		s.replaceChain(newBlockchain)
		spew.Dump(s.chain)

		return newBlock, nil
	}

	return Block{}, errors.New("Failed to add new block for payload")
}

// create a new block using previous block's hash
func (s *BlockchainService) generateBlock(oldBlock *Block, payload int) Block {
	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Payload = payload
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(&newBlock)

	return newBlock
}

// make sure the chain we're checking is longer than the current blockchain
func (s *BlockchainService) replaceChain(newChain []*Block) {
	if len(newChain) > len(s.chain) {
		s.chain = newChain
	}
}

// Chain returns the whole chain
func (s *BlockchainService) Chain() []*Block {
	return s.chain
}

func newBlockchain() []*Block {
	chain := []*Block{}

	t := time.Now()
	genesisBlock := Block{0, t.String(), 0, "Genesis-Block", ""}
	spew.Dump(genesisBlock)

	return append(chain, &genesisBlock)
}

// make sure block is valid by checking index, and comparing the hash of the previous block
func isBlockValid(newBlock *Block, oldBlock *Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

// SHA256 hasing
func calculateHash(block *Block) string {
	record := string(block.Index) + block.Timestamp + string(block.Payload) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)

	return hex.EncodeToString(hashed)
}

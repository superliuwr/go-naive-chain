package service

import (
	"fmt"

	"github.com/superliuwr/go-naive-chain/lib/util"
)

// Miner defines a miner service
type Miner struct {
}

// NewMiner returns a new instance of Miner
func NewMiner() *Miner {
	return &Miner{}
}

// ProofOfWork returns the solution to the PoW problem
func (m *Miner) ProofOfWork(lastProof int) int {
	proof := 0

	for !m.validateProof(lastProof, proof) {
		proof++
	}

	return proof
}

func (m *Miner) validateProof(lastProof int, proof int) bool {
	guess := fmt.Sprintf("%d%d", lastProof, proof)
	guessHash := util.HashSha256([]byte(guess))

	return guessHash[:4] == "0000"
}

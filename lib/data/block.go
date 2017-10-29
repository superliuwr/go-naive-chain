package data

import (
	"time"
)

// Block defines a block in the blockchain
type Block struct {
	// Index is a sequence number of this block
	Index int
	// PreviousHash is the hash value of the previous block
	PreviousHash string
	// Proof is the PoW(Proof of Work) of mining
	Proof        int
	Timestamp    time.Time
	Transactions []Transaction
}

// Transaction defines a transaction
type Transaction struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Amount    int    `json:"amount"`
}

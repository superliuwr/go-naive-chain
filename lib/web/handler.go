package web

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/superliuwr/go-naive-chain/lib/data"
	"github.com/superliuwr/go-naive-chain/lib/service"
)

type handler struct {
	blockchain *service.Blockchain
	nodeID     string
}

type response struct {
	value      interface{}
	statusCode int
	err        error
}

func newHandler(blockchain *service.Blockchain, nodeID string) http.Handler {
	h := handler{blockchain, nodeID}

	mux := http.NewServeMux()
	mux.HandleFunc("/transactions/new", buildResponse(h.AddTransaction))

	return mux
}

func buildResponse(h func(io.Writer, *http.Request) response) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := h(w, r)

		msg := resp.value
		if resp.err != nil {
			msg = resp.err.Error()
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(resp.statusCode)
		if err := json.NewEncoder(w).Encode(msg); err != nil {
			log.Printf("could not encode response to output: %v", err)
		}
	}
}

// AddTransaction returns a handler for handling adding transaction requests
func (h *handler) AddTransaction(w io.Writer, r *http.Request) response {
	if r.Method != http.MethodPost {
		return response{
			value:      nil,
			statusCode: http.StatusMethodNotAllowed,
			err:        fmt.Errorf("method %s not allowd", r.Method),
		}
	}

	log.Println("Adding transaction to the blockchain...")

	var tx data.Transaction
	err := json.NewDecoder(r.Body).Decode(&tx)
	index, err := h.blockchain.AddTransaction(tx)

	resp := map[string]string{
		"message": fmt.Sprintf("Transaction will be added to Block %d", index),
	}

	status := http.StatusCreated
	if err != nil {
		status = http.StatusInternalServerError
		log.Printf("there was an error when trying to add a transaction %v\n", err)
		err = fmt.Errorf("fail to add transaction to the blockchain")
	}

	return response{resp, status, err}
}

// Mine returns a handler for handling mining requests
func (h *handler) Mine(w io.Writer, r *http.Request) response {
	if r.Method != http.MethodGet {
		return response{
			value:      nil,
			statusCode: http.StatusMethodNotAllowed,
			err:        fmt.Errorf("method %s not allowd", r.Method),
		}
	}

	log.Println("Starting mining")

	lastBlock, err := h.blockchain.LastBlock()
	if err != nil {
		status := http.StatusInternalServerError
		log.Printf("there was an error when trying to mine %v\n", err)
		err = fmt.Errorf("fail to mine")

		return response{nil, status, err}
	}

	lastProof := lastBlock.Proof
	miner := service.NewMiner()

	proof := miner.ProofOfWork(lastProof)

	newTX := data.Transaction{Sender: "system", Recipient: h.nodeID, Amount: 1}
	_, err = h.blockchain.AddTransaction(newTX)
	if err != nil {
		status := http.StatusInternalServerError
		log.Printf("there was an error when trying to mine %v\n", err)
		err = fmt.Errorf("fail to mine")

		return response{nil, status, err}
	}

	// TODO calculate hash of last block
	block, err := h.blockchain.AddBlock(proof, "")
	if err != nil {
		status := http.StatusInternalServerError
		log.Printf("there was an error when trying to mine %v\n", err)
		err = fmt.Errorf("fail to mine")

		return response{nil, status, err}
	}

	resp := map[string]interface{}{"message": "New Block Forged", "block": block}

	return response{resp, http.StatusOK, nil}
}

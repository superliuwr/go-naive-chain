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

func (h *handler) AddTransaction(w io.Writer, r *http.Request) response {
	if r.Method != http.MethodPost {
		return response{
			value:      nil,
			statusCode: http.StatusMethodNotAllowed,
			err:        fmt.Errorf("method %s not allowd", r.Method),
		}
	}

	log.Printf("Adding transaction to the blockchain...\n")

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

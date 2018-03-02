package web

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/superliuwr/go-naive-chain/lib/service"
)

// Message takes incoming JSON payload for writing heart rate
type Message struct {
	Payload int
}

// Server defines a Web server
type Server struct {
	port              string
	blockchainService *service.BlockchainService
}

// NewServer returns a new Server instance
func NewServer(port string, blockchainService *service.BlockchainService) *Server {
	return &Server{
		port:              port,
		blockchainService: blockchainService,
	}
}

// Start starts a web server
func (s *Server) Start() error {
	mux := s.makeMuxRouter()

	log.Println("Listening for HTTP connections on ", s.port)

	httpServer := &http.Server{
		Addr:           ":" + s.port,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := httpServer.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

// create handlers
func (s *Server) makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()

	muxRouter.HandleFunc("/", s.handleGetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/", s.handleWriteBlock).Methods("POST")

	return muxRouter
}

// write blockchain when we receive an http request
func (s *Server) handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	chains := s.blockchainService.Chain()
	bytes, err := json.MarshalIndent(chains, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	io.WriteString(w, string(bytes))
}

// takes JSON payload as an input for heart rate (BPM)
func (s *Server) handleWriteBlock(w http.ResponseWriter, r *http.Request) {
	var m Message

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	newBlock, err := s.blockchainService.Add(m.Payload)
	if err != nil {
		respondWithJSON(w, r, http.StatusInternalServerError, m)
		return
	}

	respondWithJSON(w, r, http.StatusCreated, newBlock)
}

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}

	w.WriteHeader(code)
	w.Write(response)
}

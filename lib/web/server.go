package web

import (
    "fmt"
    "log"
    "net/http"

    "github.com/superliuwr/go-naive-chain/lib/service"    
)

// Server defines a Web server
type Server struct {
    Port *string
    Blockchain *service.Blockchain
    NodeID string
}

// NewServer returns a new Server instance
func NewServer(port *string, blockchain *service.Blockchain, nodeID string) *Server {
    return &Server{
        Port: port,
        Blockchain: blockchain,
        NodeID: nodeID,
    }
}

// Start starts a web server
func (s *Server) Start() error {
    // Configure HTTP routes
    http.Handle("/", newHandler(s.Blockchain, s.NodeID))

    // Start HTTP server
    log.Printf("Starting server on port:8080\n")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        return fmt.Errorf("ListenAndServe: %s", err.Error())
    }

    return nil
}
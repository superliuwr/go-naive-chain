package web

import (
    "net/http"
	"log"
	
    "github.com/superliuwr/go-naive-chain/lib/util"
)

// Start starts an HTTP server
func Start() {
    // Configure HTTP routes
    http.Handle("/api/", util.Log(APIMux))

    // Start HTTP server
    log.Printf("Starting server on port:8080\n")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
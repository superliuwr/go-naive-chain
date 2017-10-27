package util

import (
  "crypto/sha256"
  "encoding/json"
  "fmt"  
  "log"
  "net/http"
  "time"

  "github.com/superliuwr/go-naive-chain/lib/data"    
)

// Log returns a wrapped HTTP handler with which every request/response is logged
func Log(handler http.Handler) http.Handler {
  return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
    start := time.Now()
    handler.ServeHTTP(res, req)
    end := time.Since(start)
    log.Printf("%s %s %s %s", req.Host, req.URL, req.Method, end)
  })
}

// Hash returns a string of hash value representing the given Block
func Hash(block data.Block) ([]byte, error) {
  bytes, err := json.Marshal(block)
  if err != nil {
    return nil, fmt.Errorf("unable to generate hash: %s", err.Error())
  }

  sha := sha256.New()
  _, err = sha.Write(bytes)
  if err != nil {
    return nil, fmt.Errorf("unable to generate hash: %s", err.Error())
  }

  return sha.Sum(nil), nil
}
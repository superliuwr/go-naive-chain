package util

import (
  "crypto/sha256"
  "fmt"  
  "log"
  "net/http"
  "time"
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

// HashSha256 returns a hash string of the given bytes using SHA256 algorithm
func HashSha256(bytes []byte) string {
  return fmt.Sprintf("%x", sha256.Sum256(bytes))
}
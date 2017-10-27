package web

import (
    "fmt"
    "net/http"
)

// APIMux defines an HTTP server mux
var APIMux = http.NewServeMux()

// Setup Routes with Mux
func init() {
  fmt.Printf("Init APIMux\n")
  APIMux.HandleFunc("/api/test", test)
  APIMux.HandleFunc("/api/other", other)
}


// Handle /api/test
func test(res http.ResponseWriter, req *http.Request) {
  if req.Method != "GET" {
    http.Error(res, http.StatusText(405), 405)
    return
  }
  fmt.Fprintf(res, "Test Passed!")
}

// Handle /api/other
func other(res http.ResponseWriter, req *http.Request) {
  if req.Method != "GET" {
    http.Error(res, http.StatusText(405), 405)
    return
  }
  fmt.Fprintf(res, "Other Test Passed!")
}
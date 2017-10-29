package util

import (
  "crypto/sha256"
  "fmt"
)

// HashSha256 returns a hash string of the given bytes using SHA256 algorithm
func HashSha256(bytes []byte) string {
  return fmt.Sprintf("%x", sha256.Sum256(bytes))
}
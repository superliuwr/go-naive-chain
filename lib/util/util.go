package util

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

// HashSha256 returns a hash string of the given bytes using SHA256 algorithm
func HashSha256(bytes []byte) string {
	return fmt.Sprintf("%x", sha256.Sum256(bytes))
}

// PseudoUUID returns a pseudo random string
func PseudoUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return ""
	}

	return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

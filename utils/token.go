package utils

import (
	"crypto/sha256"
	"crypto/rand"
  "fmt"
)

func generateToken() string {
	randomByte := make([]byte, 10)
	rand.Read(randomByte)
	return fmt.Sprintf("%x", sha256.Sum256(randomByte))
}
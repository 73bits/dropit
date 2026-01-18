package util

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateID() string {
	b := make([]byte, 6)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}

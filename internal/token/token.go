package token

import (
	"math/rand"
	"time"
)

var (
	charset    string     = "abcdefghijklmnopqrstuvwxyz"
	seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func New(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

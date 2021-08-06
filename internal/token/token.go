package token

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

var (
	charset    string     = "abcdefghijklmnopqrstuvwxyz"
	seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
)

type Handler struct {
	mu      sync.Mutex
	fp      *os.File
	wg      *sync.WaitGroup
	toWrite chan string
	toErr   chan error
}

// Create creates <amount> tokens of length <lenght> and save it to a file
func (h *Handler) Create(length, amount int) (*sync.WaitGroup, chan error) {
	chunk := ""
	chunkCounter := 0
	chunkLength := 100
	for i := 0; i < amount; i++ {
		if chunkCounter == chunkLength {
			h.toWrite <- chunk
			chunk = ""
			chunkCounter = 0
		}
		chunk += fmt.Sprintf("%s\n", h.GetToken(length))
		chunkCounter++
	}
	if chunk != "" {
		h.toWrite <- chunk
	}

	close(h.toWrite)
	return h.wg, h.toErr
}

func (h *Handler) GetToken(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func (h *Handler) SaveChunk(chunk string) {
	h.toWrite <- chunk
}

func New(filePath string) (*Handler, error) {
	fp, err := os.Create(filePath)
	if err != nil {
		return &Handler{}, err
	}
	return &Handler{
		mu:      sync.Mutex{},
		fp:      fp,
		toWrite: make(chan string),
	}, nil
}

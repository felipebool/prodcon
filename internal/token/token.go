package token

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Handler struct {
	fp *os.File
}

// Create creates <amount> tokens of length <lenght> and save it to a file
func (h *Handler) Create(length, amount int) error {
	chunk := ""
	chunkCounter := 0
	chunkLength := 100
	for i := 0; i < amount; i++ {
		if chunkCounter == chunkLength {
			if _, err := h.fp.WriteString(chunk); err != nil {
				return err
			}
			chunk = ""
			chunkCounter = 0
			continue
		}
		chunk += fmt.Sprintf("%s\n", h.GetToken(length))
		chunkCounter++
	}
	if _, err := h.fp.WriteString(chunk); err != nil {
		return err
	}
	return nil
}

func (h *Handler) GetToken(length int) string {
	var charset string = "abcdefghijklmnopqrstuvwxyz"
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func New(filePath string) (*Handler, error) {
	fp, err := os.Create(filePath)
	if err != nil {
		return &Handler{}, err
	}
	return &Handler{fp: fp}, nil
}

package main

import (
	"flag"
	"log"

	"github.com/felipebool/prodcon/internal/token"
)

var length = flag.Int("length", 7, "token length")
var amount = flag.Int("amount", 1_000_000, "number of tokens to create")
var path = flag.String("path", "storage/tokens", "save tokens to file")

// run creates, and saves, tokens into file located in the value passed to filePath
func run(tokenAmount, tokenLength int, filePath string) error {
	handler, err := token.New(filePath)
	if err != nil {
		return err
	}
	return handler.Create(tokenLength, tokenAmount)
}

func main() {
	flag.Parse()
	if err := run(*amount, *length, *path); err != nil {
		log.Fatal(err)
	}
}

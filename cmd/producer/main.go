package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/felipebool/prodcon/internal/token"
)

//var amount = flag.Int("amount", 10_000_000, "number of tokens to create")
var length = flag.Int("length", 7, "token length")
var amount = flag.Int("amount", 10, "number of tokens to create")
var path = flag.String("path", "storage/tokens", "save tokens to file")

func run(tokenAmount, tokenLength int, filePath string) error {
	fp, err := os.Create(filePath)
	if err != nil {
		return err
	}

	for i := 0; i < tokenAmount; i++ {
		_, err := fp.WriteString(fmt.Sprintf("%s\n", token.New(tokenLength)))
		if err != nil {
			return err
		}
	}
	return fp.Close()
}

func main() {
	flag.Parse()
	log.Fatal(run(*amount, *length, *path))
}

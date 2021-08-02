package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/felipebool/prodcon/internal/cache"
)

var path = flag.String("path", "storage/tokens", "read tokens from file")

func run(c *cache.Cache, filePath string) error {
	fp, err := os.Open(filePath)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return fp.Close()
}

func main() {
	flag.Parse()
	log.Fatal(run(cache.New(), *path))
}

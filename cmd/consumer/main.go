package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/felipebool/prodcon/internal/cache"
	"github.com/felipebool/prodcon/internal/token"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

var path = flag.String("path", "storage/tokens", "read tokens from file")

func populate(c *cache.Cache, db *sqlx.DB, filePath string) error {
	if err := warmUpCache(c, filePath); err != nil {
		return err
	}

	// populate DB
	if err := populateDatabase(c, db, 10); err != nil {
		return err
	}

	return nil
}

func warmUpCache(c *cache.Cache, filePath string) error {
	fp, err := os.Open(filePath)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		c.Save(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return fp.Close()
}

func populateDatabase(c *cache.Cache, db *sqlx.DB, workers int) error {
	tokens := make(chan token.Entry, 100)

	wg := &sync.WaitGroup{}
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for t := range tokens {
				stmt := fmt.Sprintf(
					"INSERT INTO token (value, total) VALUES ('%s', %d)",
					t.Value,
					t.Total,
				)
				db.Exec(stmt)
			}
		}()
	}

	for value, total := range c.Entries {
		tokens <- token.Entry{Value: value, Total: total}
	}
	close(tokens)
	wg.Wait()
	return nil
}

func run(c *cache.Cache, db *sqlx.DB, filePath string) error {
	if err := populate(c, db, filePath); err != nil {
		return err
	}
	return nil
}

func main() {
	flag.Parse()

	dsn := "postgres://user@localhost:5433/prodcon?sslmode=disable"
	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := run(cache.New(), db, *path); err != nil {
		log.Fatal(err)
	}
}

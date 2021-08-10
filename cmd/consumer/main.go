package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/felipebool/prodcon/internal/cache"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

var path = flag.String("path", "storage/tokens", "read tokens from file")
var batchSize = flag.Int("batch", 100, "batch size for insertions")
var workersSize = flag.Int("workers", 30, "number of workers to access db")

// populate calls warmUpCache, generateReport and populateDatabase
func populate(c *cache.Cache, db *sqlx.DB, fp io.Reader, workers, batch int) error {
	wg := &sync.WaitGroup{}

	total, err := warmUpCache(c, fp)
	if err != nil {
		return err
	}

	wg.Add(1)
	go generateReport(c, total, wg)

	if err := populateDatabase(c, db, workers, batch); err != nil {
		return err
	}

	wg.Wait()
	return nil
}

// warmUpCache reads the entries from file and savem them to memory
func warmUpCache(c *cache.Cache, fp io.Reader) (int, error) {
	amount := 0
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		amount++
		c.Save(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return amount, nil
}

// populateDatabase reads batches of entries from channel and save them into db
func populateDatabase(c *cache.Cache, db *sqlx.DB, workers, batch int) error {
	tokens := make(chan string, 1000)
	wg := &sync.WaitGroup{}
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for values := range tokens {
				stmt := fmt.Sprintf(
					"INSERT INTO token (value, total) VALUES %s",
					values,
				)
				if _, err := db.Exec(stmt); err != nil {
					fmt.Println(err)
				}
			}
		}()
	}

	// produce tokens
	entriesCount := 0
	insertValues := ""
	for value, total := range c.Entries {
		if entriesCount == batch {
			tokens <- insertValues[:len(insertValues)-1]
			entriesCount = 0
			insertValues = ""
		}
		insertValues += fmt.Sprintf("('%s', %d),", value, total)
		entriesCount++
	}
	tokens <- insertValues[:len(insertValues)-1]
	close(tokens)
	wg.Wait()
	return nil
}

// generateReport produces a list of all non-unique tokens and their frequencies
func generateReport(c *cache.Cache, tokenAmount int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("token\tfrequency (%d)\n", tokenAmount)
	for value, total := range c.Entries {
		if total > 1 {
			fmt.Printf("%s\t%f\n", value, float32(total)/float32(tokenAmount))
		}
	}
}

func run(c *cache.Cache, db *sqlx.DB, filePath string, batch, workers int) error {
	fp, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer fp.Close()

	if err := populate(c, db, fp, batch, workers); err != nil {
		return err
	}
	return nil
}

func main() {
	flag.Parse()
	schema := `CREATE TABLE IF NOT EXISTS token (
		value varchar(7),
		total integer
	);`

	dsn := "postgres://user@localhost:5433/prodcon?sslmode=disable"
	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// create database schema
	if _, err = db.Exec(schema); err != nil {
		log.Fatal(err)
	}

	// run application
	if err := run(cache.New(), db, *path, *batchSize, *workersSize); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"bufio"
	"database/sql"
	"flag"
	"log"
	"os"

	"github.com/felipebool/prodcon/internal/cache"
	_ "github.com/go-sql-driver/mysql"
)

var path = flag.String("path", "storage/tokens", "read tokens from file")

func populate(c *cache.Cache, db *sql.DB, filePath string) error {
	if err := warmUpCache(c, filePath); err != nil {
		return err
	}

	// create a pool of workers

	// populate DB
	for key, value := range c.Entries {
		stmt, err := db.Prepare("INSERT INTO token(key, value) VALUES(?)")
		if err != nil {
			return err
		}
		_, err = stmt.Exec(key, value)
		if err != nil {
			return err
		}
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

func run(c *cache.Cache, db *sql.DB, filePath string) error {
	if err := populate(c, db, filePath); err != nil {
		return err
	}
	return nil
}

func main() {
	flag.Parse()

	dsn := ""
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	log.Fatal(run(cache.New(), db, *path))
}

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
	fp, err := os.Open(filePath)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		if err := writeBack(c, db, scanner.Text()); err != nil {
			return err
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return fp.Close()
}

func writeBack(c *cache.Cache, db *sql.DB, token string) error {
	// write value to cache
	c.Save(token)

	// check cache
	if c.Fetch(token) == 1 {
		c.Save(token)
		// save token into database
		stmt, err := db.Prepare("INSERT INTO token(value) VALUES(?)")
		if err != nil {
			return err
		}

		res, err := stmt.Exec(token)
		if err != nil {
			return err
		}

		_, err = res.LastInsertId()
		if err != nil {
			return err
		}
	}

	return nil
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

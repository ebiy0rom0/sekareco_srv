package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sekareco_srv/env"
	"sekareco_srv/util"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tanimutomo/sqlfile"
)

func main() {
	var setup bool
	flag.BoolVar(&setup, "setup", true, "exec setup or cleanup")
	flag.Parse()

	if setup {
		Setup()
	} else {
		Cleanup()
	}
}

// testing before setup package
func Setup() {
	fmt.Println("---------- start testing setup. ----------")

	source := filepath.Join(util.RootDir(), env.DbDir, env.DbFile)

	// Cleaning DB file if left before testing DB
	if _, err := os.Stat(source); err == nil {
		if err := os.Remove(source); err != nil {
			log.Fatalf("db cleaning error: %s", err.Error())
		}

	}

	// DB file automatically create
	db, err := sql.Open("sqlite3", source)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	// create base table and insert test data
	s := sqlfile.New()
	err = s.Directory(util.RootDir() + "/docs/db/")
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	_, err = s.Exec(db)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	err = s.Directory(util.RootDir() + "/test/data/")
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	_, err = s.Exec(db)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	fmt.Println("---------- complete testing setup. ----------")
}

// testing after cleaning package
func Cleanup() {
	fmt.Println("---------- start testing cleanup. ----------")
	// something todo...
	fmt.Println("---------- complete testing cleanup. ----------")
}

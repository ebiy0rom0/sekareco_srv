package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sekareco_srv/infra"
	"sekareco_srv/util"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tanimutomo/sqlfile"
)

// testing before setup package
func main() {
	fmt.Println("---------- start testing setup. ----------")

	if err := infra.LoadEnv(".env.testing"); err != nil {
		log.Fatalf("env load error: %s\n", err.Error())
	}

	dbPath := util.RootDir() + os.Getenv("DB_PATH") + os.Getenv("DB_NAME")

	// Cleaning DB file if left before testing DB
	if _, err := os.Stat(dbPath); err == nil {
		if err := os.Remove(dbPath); err != nil {
			log.Fatalf("db cleaning error: %s", err.Error())
		}

	}

	// DB file automatically create
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	// insert test data
	s := sqlfile.New()
	err = s.Directory(util.RootDir() + "/test/data/")
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	s.Exec(db)

	fmt.Println("---------- complete testing setup. ----------")
}

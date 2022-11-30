package main

import (
	"log"
	"os"
	"path/filepath"
	"sekareco_srv/infra"
	"sekareco_srv/infra/sql"
	"sekareco_srv/util"

	"github.com/tanimutomo/sqlfile"
)

func main() {
	if err := initialize(); err != nil {
		log.Fatalf("%+v", err)
	}
}

func initialize() error {
	if err := infra.LoadEnv("dev"); err != nil {
		return nil
	}

	if err := makeDirectories(); err != nil {
		return err
	}

	if err := storeTestRecords(); err != nil {
		return err
	}
	return nil
}

// makeDir is make the necessary directries.
func makeDirectories() error {
	directories := []string{"bin", "db", "log", "coverage"}
	for _, dir := range directories {
		dir = filepath.Join(util.RootDir(), dir)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err := os.Mkdir(dir, 0755); err != nil {
				return err
			}
		}
	}
	return nil
}

func storeTestRecords() error {
	name := os.Getenv("DB_NAME")
	con, err := sql.NewConnection("", "", "", name)
	if err != nil {
		return err
	}

	s := sqlfile.New()

	// TODO: exec sql/ directory, not single file.
	file := filepath.Join(util.RootDir(), "tools/initializer/sql/master.sql")
	if err := s.File(file); err != nil {
		return err
	}

	if _, err := s.Exec(con); err != nil {
		return err
	}
	return nil
}

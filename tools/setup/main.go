package main

import (
	"log"
	"os"
	"path/filepath"
	"sekareco_srv/env"
	"sekareco_srv/infra/sql"
	"sekareco_srv/util"

	"github.com/ebiy0rom0/errors"

	"github.com/tanimutomo/sqlfile"
)

func main() {
	if err := initialize(); err != nil {
		log.Fatalf("%+v", err)
	}
}

func initialize() error {
	if err := makeDirectories(); err != nil {
		return errors.Wrap(err, "failed to make directories")
	}

	if err := storeTestRecords(); err != nil {
		return errors.Wrap(err, "failed to store testdata")
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
				return errors.New(err.Error())
			}
		}
	}
	return nil
}

func storeTestRecords() error {
	con, err := sql.NewConnection("", "", "", env.DbFile)
	if err != nil {
		return errors.New(err.Error())
	}

	s := sqlfile.New()

	// TODO: exec sql/ directory, not single file.
	file := filepath.Join(util.RootDir(), "tools/setup/sql/master.sql")
	if err := s.File(file); err != nil {
		return errors.New(err.Error())
	}

	if _, err := s.Exec(con.DB); err != nil {
		return errors.New(err.Error())
	}
	return nil
}

package sql

import (
	"database/sql"
	"net"
	"time"

	"github.com/ebiy0rom0/errors"

	"github.com/go-sql-driver/mysql"
)

// openMysql establishes a connection with db for MySQL
// and returns a pointer to the connection.
func openMysql(user, pass, host, schema string) (*sql.DB, error) {
	//user:password@tcp(host:port)/dbname
	// source := fmt.Sprintf(
	// 	"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true",
	// 	user, pass, host, schema,
	// )

	// tcp connection test
	if _, err := net.DialTimeout("tcp", host, 1*time.Second); err != nil {
		return nil, err
	}

	c := mysql.Config{
		User:   user,
		Passwd: pass,
		Addr:   host,
		DBName: schema,
	}

	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return db, nil
}

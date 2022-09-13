package sql

import (
	"database/sql"
	"fmt"
)

// openMysql establishes a connection with db for MySQL
// and returns a pointer to the connection.
func openMysql(user, pass, host, schema string) (*sql.DB, error) {
	//user:password@tcp(host:port)/dbname
	source := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true",
		user, pass, host, schema,
	)

	return sql.Open("mysql", source)
}
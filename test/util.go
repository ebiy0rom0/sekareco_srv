package test

import (
	"log"
	"path/filepath"
	"sekareco_srv/domain/infra"
	"sekareco_srv/env"
	"sekareco_srv/infra/middleware"
	"sekareco_srv/infra/sql"
	infraIf "sekareco_srv/interface/infra"
	"sekareco_srv/util"

	"github.com/jmoiron/sqlx"
)

var sqlHandler infraIf.SqlHandler
var txHandler infraIf.TxHandler

var authMiddleware *middleware.AuthMiddleware

func Initialize() {
	source := filepath.Join(util.RootDir(), env.DbDir, env.DbFile)

	con, err := sqlx.Connect("sqlite3", source)
	if err != nil {
		log.Fatalf("Failed connect db: %+v\n", err)
	}

	sqlHandler = sql.NewSqlHandler(con)
	txHandler = sql.NewTxHandler(con)

	am := middleware.NewAuthMiddleware()
	authMiddleware = am
}

func InjectSqlHandler() infraIf.SqlHandler {
	return sqlHandler
}

func InjectTxHandler() infraIf.TxHandler {
	return txHandler
}

func InjectAuthMiddleware() *middleware.AuthMiddleware {
	return authMiddleware
}

func InjectTokenManager() infra.TokenManager {
	return authMiddleware
}

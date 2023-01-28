package test

import (
	"log"
	infraDoamin "sekareco_srv/domain/infra"
	"sekareco_srv/env"
	"sekareco_srv/infra"
	"sekareco_srv/infra/middleware"
	"sekareco_srv/infra/sql"
	infraIf "sekareco_srv/interface/infra"
)

var sqlHandler infraIf.SqlHandler
var txHandler infraIf.TxHandler

var authMiddleware *middleware.AuthMiddleware

func Setup() {
	if err := infra.LoadEnv("test"); err != nil {
		log.Fatalf("env load error: %s\n", err.Error())
	}

	con, err := sql.NewConnection("", "", "", env.DbFile)
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

func InjectTokenManager() infraDoamin.TokenManager {
	return authMiddleware
}

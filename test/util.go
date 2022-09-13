package test

import (
	"log"
	"os"
	infra__ "sekareco_srv/domain/infra"
	"sekareco_srv/infra"
	"sekareco_srv/infra/middleware"
	"sekareco_srv/infra/sql"
	infra_ "sekareco_srv/interface/infra"
)

var sqlHandler infra_.SqlHandler
var txHandler infra_.TxHandler

var authMiddleware *middleware.AuthMiddleware

func Setup() {
	if err := infra.LoadEnv("test"); err != nil {
		log.Fatalf("env load error: %s\n", err.Error())
	}

	con, err := sql.NewConnection("", "", "", os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatalf("Failed connect db: %+v\n", err)
	}

	sqlHandler = sql.NewSqlHandler(con)
	txHandler = sql.NewTxHandler(con)

	am := middleware.NewAuthMiddleware()
	authMiddleware = am
}

func InjectSqlHandler() infra_.SqlHandler {
	return sqlHandler
}

func InjectTxHandler() infra_.TxHandler {
	return txHandler
}

func InjectAuthMiddleware() *middleware.AuthMiddleware {
	return authMiddleware
}

func InjectTokenManager() infra__.TokenManager {
	return authMiddleware
}

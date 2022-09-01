package test

import (
	"log"
	"os"
	infra__ "sekareco_srv/domain/infra"
	"sekareco_srv/infra"
	"sekareco_srv/infra/middleware"
	sql_ "sekareco_srv/infra/sql"
	infra_ "sekareco_srv/interface/infra"
	"sekareco_srv/util"
)

var sqlHandler infra_.SqlHandler
var txHandler infra_.TxHandler

var authMiddleware *middleware.AuthMiddleware

func Setup() {
	if err := infra.LoadEnv("test"); err != nil {
		log.Fatalf("env load error: %s\n", err.Error())
	}

	dbPath := util.RootDir() + os.Getenv("DB_PATH") + os.Getenv("DB_NAME")
	sh, th, err := sql_.NewSqlHandler(dbPath)
	if err != nil {
		log.Fatalf("db connection error: %s\n", err.Error())
	}

	sqlHandler = sh
	txHandler = th

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

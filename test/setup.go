package test

import (
	"log"
	"os"
	infra__ "sekareco_srv/domain/infra"
	"sekareco_srv/infra"
	"sekareco_srv/infra/middleware"
	"sekareco_srv/infra/sql"
	infra_ "sekareco_srv/interface/infra"
	"sekareco_srv/util"
)

var setup bool = false

var sqlHandler infra_.SqlHandler
var txHandler infra_.TxHandler

var authMiddleware *middleware.AuthMiddleware

func Setup() {
	// if already setup done, don't re-exec setup
	if setup {
		return
	}
	if err := infra.LoadEnv(".env.testing"); err != nil {
		log.Fatalf("env load error: %s\n", err.Error())
	}

	dbPath := util.RootDir() + os.Getenv("DB_PATH") + os.Getenv("DB_NAME")
	sh, th, err := sql.NewSqlHandler(dbPath)
	if err != nil {
		log.Fatalf("db setup error: %s\n", err.Error())
	}

	am := middleware.NewAuthMiddleware()

	sqlHandler = sh
	txHandler = th
	authMiddleware = am

	// re-exec setup control
	setup = true
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

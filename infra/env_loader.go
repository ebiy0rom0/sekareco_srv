package infra

import (
	"sekareco_srv/util"

	"github.com/joho/godotenv"
)

const ENV_FILE_PATH = "/config/"

func LoadEnv(target string) error {
	return godotenv.Load(util.RootDir() + ENV_FILE_PATH + target)
}

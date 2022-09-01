package infra

import (
	"fmt"
	"sekareco_srv/util"

	"github.com/joho/godotenv"
)

func LoadEnv(env string) error {
	return godotenv.Load(fmt.Sprintf("%s/env/%s.env", util.RootDir(), env))
}

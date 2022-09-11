package infra

import (
	"fmt"
	"sekareco_srv/util"

	"github.com/joho/godotenv"
)

// LoadEnv loads enviroment valiables from the .env file in the {rootDir}/env/ directory.
// Switching the .env file to be loaded by changing the argument stage.
func LoadEnv(stage string) error {
	return godotenv.Load(fmt.Sprintf("%s/env/%s.env", util.RootDir(), stage))
}

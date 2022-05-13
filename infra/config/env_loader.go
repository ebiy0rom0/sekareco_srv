package loader

import "github.com/joho/godotenv"

const ENV_FILE_PATH = "./../config/"

func LoadEnv(target string) error {
	return godotenv.Load(ENV_FILE_PATH + target)
}

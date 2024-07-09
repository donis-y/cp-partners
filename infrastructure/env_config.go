package infrastructure

import (
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv() error {
	return godotenv.Load()
}

func GetEnv(key string) string {
	return os.Getenv(key)
}

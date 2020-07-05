package env

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

// GetEnv returns the ENV var
func GetEnv(key string) string {
	return os.Getenv(key)
}

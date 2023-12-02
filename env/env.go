package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Getenv(key string) string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	return os.Getenv(key)
}

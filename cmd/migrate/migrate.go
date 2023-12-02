package main

import (
	"log"
	"os"

	"github.com/alvinfebriando/gin-gorm-skeleton/migration"
	"github.com/alvinfebriando/gin-gorm-skeleton/repository"
)

func main() {
	_ = os.Setenv("APP_ENV", "debug")

	db, err := repository.GetConnection()
	if err != nil {
		log.Println(err)
	}

	migration.Migrate(db)
}

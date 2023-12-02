package main

import (
	"os"

	"github.com/alvinfebriando/gin-gorm-skeleton/applogger"
	"github.com/alvinfebriando/gin-gorm-skeleton/migration"
	"github.com/alvinfebriando/gin-gorm-skeleton/repository"
)

func main() {
	_ = os.Setenv("APP_ENV", "debug")

	db, err := repository.GetConnection()
	if err != nil {
		applogger.Log.Error(err)
	}

	migration.Seed(db)
}

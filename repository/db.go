package repository

import (
	"os"

	"github.com/alvinfebriando/gin-gorm-skeleton/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetConnection() (*gorm.DB, error) {
	dsn := env.GetDsn()
	level := logger.Silent

	if os.Getenv("APP_ENV") == "dev" {
		level = logger.Error
	}
	if os.Getenv("APP_ENV") == "debug" {
		level = logger.Info
	}

	config := &gorm.Config{
		Logger: logger.Default.LogMode(level),
	}

	return gorm.Open(postgres.Open(dsn), config)
}

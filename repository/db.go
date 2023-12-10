package repository

import (
	"fmt"

	"github.com/alvinfebriando/gin-gorm-skeleton/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func getDsn() string {
	host := config.New().DbHost
	port := config.New().DbPort
	user := config.New().DbUser
	pass := config.New().DbPass
	dbName := config.New().DbName
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, user, pass, dbName)
	return dsn
}

func GetConnection() (*gorm.DB, error) {
	dsn := getDsn()
	level := logger.Silent

	if config.New().IsInDevMode() {
		level = logger.Error

	}
	if config.New().IsInDebugMode() {
		level = logger.Info
	}

	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(level),
	})
}

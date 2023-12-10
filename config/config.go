package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Option struct {
	AppName                     string
	AppEnv                      string
	AppHost                     string
	AppPort                     string
	DbHost                      string
	DbPort                      string
	DbUser                      string
	DbPass                      string
	DbName                      string
	ResetPasswordExpiryDuration int
	JwtExpiryDuration           time.Duration
	JwtSecret                   string
	GracefulShutdownTimeout     time.Duration
	RequestTimeout              time.Duration
}

func (o *Option) IsInDevMode() bool {
	return o.AppEnv == "dev"
}

func (o *Option) IsInDebugMode() bool {
	return o.AppEnv == "debug"
}

const defaultResetPasswordExpiryDuration = 1
const defaultJwtExpiryDuration = 1
const defaultGracefulShutdownTimeout = 5
const defaultRequestTimeout = 5

var config *Option

func New() *Option {
	if config == nil {
		config = initialize()
	}
	return config
}

func initialize() *Option {
	_ = godotenv.Load()
	appName := os.Getenv("APP_NAME")
	appEnv := os.Getenv("APP_ENV")
	appHost := os.Getenv("APP_HOST")
	appPort := os.Getenv("APP_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	jwtSecret := os.Getenv("JWT_SECRET")

	resetPasswordExpiryDurationString := os.Getenv("RESET_PASSWORD_EXPIRY_DURATION")
	resetPasswordExpiryDuration, err := strconv.Atoi(resetPasswordExpiryDurationString)
	if err != nil {
		resetPasswordExpiryDuration = defaultResetPasswordExpiryDuration
	}

	jwtExpiryDurationString := os.Getenv("JWT_EXPIRY_DURATION")
	jwtExpiryDuration, err := strconv.Atoi(jwtExpiryDurationString)
	if err != nil {
		jwtExpiryDuration = defaultJwtExpiryDuration
	}

	gracefulShutdownTimeoutString := os.Getenv("GRACEFUL_SHUTDOWN_TIMEOUT")
	gracefulShutdownTimeout, err := strconv.Atoi(gracefulShutdownTimeoutString)
	if err != nil {
		gracefulShutdownTimeout = defaultGracefulShutdownTimeout
	}

	requestTimeoutString := os.Getenv("REQUEST_TIMEOUT")
	requestTimeout, err := strconv.Atoi(requestTimeoutString)
	if err != nil {
		requestTimeout = defaultRequestTimeout
	}

	return &Option{
		AppName:                     appName,
		AppEnv:                      appEnv,
		AppHost:                     appHost,
		AppPort:                     appPort,
		DbHost:                      dbHost,
		DbPort:                      dbPort,
		DbUser:                      dbUser,
		DbPass:                      dbPass,
		DbName:                      dbName,
		ResetPasswordExpiryDuration: resetPasswordExpiryDuration,
		JwtExpiryDuration:           time.Duration(jwtExpiryDuration),
		JwtSecret:                   jwtSecret,
		GracefulShutdownTimeout:     time.Duration(gracefulShutdownTimeout),
		RequestTimeout:              time.Duration(requestTimeout),
	}
}

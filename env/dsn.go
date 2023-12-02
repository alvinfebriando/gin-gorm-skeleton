package env

import "fmt"

func GetDsn() string {
	host := Getenv("DB_HOST")
	port := Getenv("DB_PORT")
	user := Getenv("DB_USER")
	pass := Getenv("DB_PASS")
	dbName := Getenv("DB_NAME")
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, user, pass, dbName)
	return dsn
}

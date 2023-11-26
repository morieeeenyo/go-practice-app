package infra

import (
	"database/sql"
	"fmt"
	"os"
)

func InitDBConnection() * sql.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	connection, err := sql.Open("mysql", dsn)
	if err != nil {
		panic("failed to connect database")
	}
	
	return connection
}
package infra

import (
	"database/sql"
	"fmt"
	// "os"
)

func InitDBConnection() * sql.DB {
	dsn := fmt.Sprintf(
		"%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		"root",
		"127.0.0.1",
		"3306",
		"go_practice_app_development",
	)
	connection, err := sql.Open("mysql", dsn)
	if err != nil {
		panic("failed to connect database")
	}
	
	return connection
}
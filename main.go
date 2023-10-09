package main

import (
	"net/http"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	e := echo.New()

	// DB接続
	dsn := fmt.Sprintf(
		"%s:%s@tcp(db:3306)/%s?charset=utf8&parseTime=true",
		// os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		// os.Getenv("DB_PORT"),
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("failed to init database: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	log.Default().Println("success to connect db!!")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
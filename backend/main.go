/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/togo-mentor/go-practice-app/router"
)

func main() {
	e := echo.New()
	e = router.InitRouter(e)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

    /* CORS(オリジン間リソース共有)の設定 */
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        /* http://localhost:3000からの接続を許可する */
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPut,
			http.MethodPost,
			http.MethodDelete,
		},
	}))

	e.Logger.Fatal(e.Start(":8080"))
}

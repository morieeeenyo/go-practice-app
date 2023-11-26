package router

import (
	"net/http"

	"github.com/labstack/echo/v4"

	handler "github.com/togo-mentor/go-practice-app/handlers"
	"github.com/togo-mentor/go-practice-app/infra"
	"github.com/togo-mentor/go-practice-app/infra/mysql"
	service "github.com/togo-mentor/go-practice-app/services"
	usecase "github.com/togo-mentor/go-practice-app/usecases"
)

func InitRouter(e *echo.Echo) *echo.Echo {
	dbConnection := infra.InitDBConnection()
	todoRepository := mysql.NewTodoRepository(dbConnection)
	todoService := service.NewTodoService(todoRepository)
	todoUsecase := usecase.NewTodoUsecase(todoService)
	g := e.Group("/todos")
	todoHandler := handler.NewTodoHandler(todoUsecase)
	{
		g.POST("", todoHandler.CreateTodo())
		g.GET("", todoHandler.FindyAllTodos())
	}

	return e
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
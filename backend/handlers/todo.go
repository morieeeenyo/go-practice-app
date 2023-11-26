package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/togo-mentor/go-practice-app/src/database/models"
	"github.com/togo-mentor/go-practice-app/usecases"
)

type TodoHandler struct {
	usecase usecase.ITodoUsecase
}

func NewTodoHandler(usecase usecase.ITodoUsecase) *TodoHandler {
	return &TodoHandler{
		usecase,
	}
}

func (h *TodoHandler) CreateTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error
		var response *models.Todo
		ctx := c.Request().Context()

		todo := &models.Todo {}
		if err = c.Bind(todo); err != nil {
      return echo.ErrBadRequest
    }
		response, err = h.usecase.CreateTodo(ctx, todo)

		if err != nil {
			log.Println(err)
			return echo.ErrInternalServerError
		}
		return c.JSON(http.StatusCreated, response)
	}
}

func (h *TodoHandler) FindyAllTodos() echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error
		var response models.TodoSlice
	
		ctx := c.Request().Context()
	
		response, err = h.usecase.FindAllTodos(ctx)
	
		if err != nil {
			log.Println(err)
			return echo.ErrInternalServerError
		}
	
		return c.JSON(http.StatusOK, response)
	}
}
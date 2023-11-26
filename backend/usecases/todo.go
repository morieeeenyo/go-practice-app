package usecase

import (
	"context"
	// "database/sql"
	// "net/http"

	"github.com/togo-mentor/go-practice-app/services"
	"github.com/togo-mentor/go-practice-app/src/database/models"
	// "github.com/mholt/binding"
	// "github.com/volatiletech/sqlboiler/boil"
	// "github.com/labstack/echo"

	_ "github.com/go-sql-driver/mysql"
)

type ITodoUsecase interface {
	FindAllTodos(ctx context.Context) (models.TodoSlice, error)
	CreateTodo(ctx context.Context, params *models.Todo) (*models.Todo, error)
}

type TodoUsecase struct {
	svc service.ITodoService
}

func NewTodoUsecase(svc service.ITodoService) ITodoUsecase  {
	return &TodoUsecase{
		svc,
	}
}

// TODOを作成
func (u TodoUsecase) CreateTodo(ctx context.Context, params *models.Todo) (*models.Todo, error) {
	todo, err := u.svc.CreateTodo(ctx, params)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (u TodoUsecase) FindAllTodos(ctx context.Context) (models.TodoSlice, error) {
	todos, err := u.svc.FindAllTodos(ctx)
	if err != nil {
		return  nil, err
	}
	return todos, nil
}
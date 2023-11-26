package repository

import (
	"context"

	"github.com/togo-mentor/go-practice-app/src/database/models"
)

// IHogeHoge represents interface of HogeHoge
type ITodoRepository interface {
	SelectAllTodos(ctx context.Context) (models.TodoSlice, error)
	CreateTodo(ctx context.Context, params *models.Todo) (*models.Todo, error)
}
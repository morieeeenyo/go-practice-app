package service

import (
	"context"

	"github.com/togo-mentor/go-practice-app/repository"
	"github.com/togo-mentor/go-practice-app/src/database/models"
)

// interface
type ITodoService interface {
	FindAllTodos(ctx context.Context) (models.TodoSlice, error)
	CreateTodo(ctx context.Context, params *models.Todo) (*models.Todo, error)
}

// struct that meets interface
type TodoService struct {
	repo repository.ITodoRepository
}

// constructor
func NewTodoService(sr repository.ITodoRepository) ITodoService {
	return &TodoService{
		repo: sr,
	}
}

// implement methods according to interface
func (ss *TodoService) FindAllTodos(ctx context.Context) (models.TodoSlice, error) {
	return ss.repo.SelectAllTodos(ctx)
}

func (ss *TodoService) CreateTodo(ctx context.Context, params *models.Todo) (*models.Todo, error) {
	return ss.repo.CreateTodo(ctx, params)
}
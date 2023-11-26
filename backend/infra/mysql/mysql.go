package mysql

import (
	"context"
	"database/sql"

	"github.com/togo-mentor/go-practice-app/repository"
	"github.com/togo-mentor/go-practice-app/src/database/models"
	"github.com/volatiletech/sqlboiler/boil"
)

type TodoRepository struct {
	DB *sql.DB
}

func NewTodoRepository(db *sql.DB) repository.ITodoRepository {
	return &TodoRepository{
		DB: db,
	}
}

func (sr *TodoRepository) SelectAllTodos(ctx context.Context) (models.TodoSlice, error) {
	// concrete DB operation
	return models.Todos().All(ctx, sr.DB)
}

func (sr *TodoRepository) CreateTodo(ctx context.Context, params *models.Todo) (*models.Todo, error) {
	Todo := &models.Todo {
		Title: params.Title,
	}
	tx, err := sr.DB.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	err = Todo.Insert(ctx, tx ,boil.Infer())
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	response := &models.Todo{
		Title: Todo.Title,
	}
	return response, nil
}
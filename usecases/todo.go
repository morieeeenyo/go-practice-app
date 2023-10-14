package usecase

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/togo-mentor/go-practice-app/src/database/models"
	"github.com/mholt/binding"
	"github.com/volatiletech/sqlboiler/boil"

	_ "github.com/go-sql-driver/mysql"
)

type TodoUsecase struct {
	ctx context.Context
	tx  *sql.Tx
}

func NewTodoUsecase(ctx context.Context, tx *sql.Tx) *TodoUsecase {
	return &TodoUsecase{
		ctx: ctx,
		tx:  tx,
	}
}

// TODO作成APIのリクエストパラメータ
type CreateTodoInput struct {
	Title string
}

// リクエストのマッピング。ポインターレシーバーにすること
func (input *CreateTodoInput) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&input.Title: "title",
	}
}

// TODO作成APIのレスポンス
type CreateTodoOutput struct {
	ID   int64    `json:"id"`
	Title string `json:"title"`
	IsCompleted  int8    `json:"isCompleted"`
}

// TODOを作成
func (u TodoUsecase) Create(input CreateTodoInput) (CreateTodoOutput, error) {
	Todo := models.Todo{
		Title: input.Title,
	}
	err := Todo.Insert(u.ctx, u.tx, boil.Infer())
	if err != nil {
		return CreateTodoOutput{}, err
	}
	// TODO: もしかしたら作成されたレコードじゃないかも。あとで直す
	output := CreateTodoOutput{
		ID:   Todo.ID,
		Title:   Todo.Title,
		IsCompleted:   Todo.IsCompleted,
	}
	return output, nil
}
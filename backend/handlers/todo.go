package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/mholt/binding"
	"github.com/togo-mentor/go-practice-app/usecases"
	"github.com/togo-mentor/go-practice-app/db"
)

type TodoHandler struct {}

func NewTodoHandler() *TodoHandler {
	return &TodoHandler{}
}

// 失敗時のレスポンスの設定
func setErrorResponse(w http.ResponseWriter, status int) {
	// レスポンスのヘッダ設定
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// 引数のステータス設定
	w.WriteHeader(status)
}

// 成功時のレスポンスの設定
func setSuccessResponse(w http.ResponseWriter, res []byte) {
	// レスポンスの内容をjsonに変換
	w.Write(res)
}

// TODO作成API
func (h *TodoHandler) Create(w http.ResponseWriter, r *http.Request) {
	var err error
	var berr binding.Errors
	var response usecase.CreateTodoOutput
	var res []byte

	// request -> CreateTodoInput型に変換
	var request usecase.CreateTodoInput
	berr = binding.Bind(r, &request)
	if berr != nil {
		log.Println(berr)
		setErrorResponse(w, http.StatusInternalServerError)
		return
	}

	ctx := context.Background()

	// dbのコネクション設定
	conn, err := db.GetDBconnection()
	if err != nil {
		log.Println(err)
		return
	}
	// コネクションのクローズ
	defer func() {
		conn.Close()
	}()

	// トランザクションの開始
	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// トランザクションのコミット or ロールバック
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	// usecaseの初期化
	TodoUsecase := usecase.NewTodoUsecase(ctx, tx)
	// TODO作成実施
	response, err = TodoUsecase.Create(request)

	// レスポンスをjsonに変換
	res, err = json.Marshal(response)
	if err != nil {
		log.Println(err)
		setErrorResponse(w, http.StatusInternalServerError)
		return
	}
	setSuccessResponse(w, res)
}

type GetTodoOutput struct {
	ID   int64    `json:"id"`
	Title string `json:"title"`
	IsCompleted  int8    `json:"isCompleted"`
}

func (h *TodoHandler) Get(w http.ResponseWriter, r *http.Request) {
	var err error
	var response usecase.GetTodoSlice
	var res []byte

	ctx := context.Background()

	// dbのコネクション設定
	conn, err := db.GetDBconnection()
	if err != nil {
		log.Println(err)
		return
	}
	// コネクションのクローズ
	defer func() {
		conn.Close()
	}()

	// トランザクションの開始
	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// トランザクションのコミット or ロールバック
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	// usecaseの初期化
	TodoUsecase := usecase.NewTodoUsecase(ctx, tx)
	response, err = TodoUsecase.Get()

	// レスポンスをjsonに変換
	res, err = json.Marshal(response)
	if err != nil {
		log.Println(err)
		setErrorResponse(w, http.StatusInternalServerError)
		return
	}
	setSuccessResponse(w, res)
}
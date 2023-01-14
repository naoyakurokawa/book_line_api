package main

import (
	"context"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/naoyakurokawa/book_line_api/auth"
	"github.com/naoyakurokawa/book_line_api/clock"
	"github.com/naoyakurokawa/book_line_api/config"
	"github.com/naoyakurokawa/book_line_api/handler"
	"github.com/naoyakurokawa/book_line_api/service"
	"github.com/naoyakurokawa/book_line_api/store"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	r := mux.NewRouter()
	clocker := clock.RealClocker{}
	v := validator.New()
	db, cleanup, err := store.New(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}
	rep := store.Repository{Clocker: clocker}

	// ユーザー登録
	ru := &handler.CreateUserHandler{
		Service:   &service.CreateUserService{DB: db, Repo: &rep},
		Validator: v,
	}
	r.HandleFunc("/register", ru.ServeHTTP).Methods(http.MethodPost)

	// ログイン
	rcli, err := store.NewKVS(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}
	jwter, err := auth.NewJWTer(rcli, clocker)
	if err != nil {
		return nil, cleanup, err
	}
	l := &handler.Login{
		Service: &service.Login{
			DB:             db,
			Repo:           &rep,
			TokenGenerator: jwter,
		},
		Validator: v,
	}
	r.HandleFunc("/login", l.ServeHTTP).Methods("POST", "OPTIONS")

	// 認証が必要なメソッドを分けるため、Subrouterを実行
	auth := r.PathPrefix("").Subrouter()
	// 本登録
	cb := &handler.CreateBooksHandler{
		Service:   &service.CreateBookService{DB: db, Repo: &rep},
		Validator: v,
	}
	auth.HandleFunc("/create/book", cb.CreateBook).Methods(http.MethodPost)
	// 本一覧取得
	fb := &handler.FetchBooksHandler{
		Service: &service.FetchBooksService{DB: db, Repo: &rep},
	}
	auth.HandleFunc("/books", fb.FetchBooks).Methods(http.MethodGet)

	// 本のメモ登録
	cbm := &handler.CreateBookMemoHandler{
		Service:   &service.CreateBookMemoService{DB: db, Repo: &rep},
		Validator: v,
	}
	auth.HandleFunc("/create/book/memo", cbm.CreateBookMemo).Methods(http.MethodPost)

	// 本のメモ一覧取得
	fbm := &handler.FetchBookMemosHandler{
		Service:   &service.FetchBookMemosService{DB: db, Repo: &rep},
		Validator: v,
	}
	auth.HandleFunc("/books/memos", fbm.ServeHTTP).Methods(http.MethodGet)

	// 認証ミドルウェア使用
	// auth.Use(handler.AuthMiddleware(jwter))
	r.Use(handler.CorsMiddlewareFunc())

	return r, cleanup, nil
}

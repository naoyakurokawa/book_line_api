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
	ru := &handler.RegisterUser{
		Service:   &service.RegisterUser{DB: db, Repo: &rep},
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
	r.HandleFunc("/login", l.ServeHTTP).Methods(http.MethodPost)

	// 認証が必要なメソッドを分けるため、Subrouterを実行
	auth := r.PathPrefix("").Subrouter()
	// 本一覧取得
	lt := &handler.ListBook{
		Service: &service.ListBook{DB: db, Repo: &rep},
	}
	auth.HandleFunc("/books", lt.ServeHTTP).Methods(http.MethodGet)

	lbm := &handler.ListBookMemo{
		Service:   &service.ListBookMemos{DB: db, Repo: &rep},
		Validator: v,
	}
	auth.HandleFunc("/books/memos", lbm.ServeHTTP).Methods(http.MethodGet)

	// 認証ミドルウェア使用
	// auth.Use(handler.AuthMiddleware(jwter))

	return r, cleanup, nil
}

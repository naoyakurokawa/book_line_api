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
	rcli, err := store.NewKVS(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}
	jwter, err := auth.NewJWTer(rcli, clocker)
	if err != nil {
		return nil, cleanup, err
	}
	lt := &handler.ListBook{
		Service: &service.ListBook{DB: db, Repo: &rep},
	}
	r.HandleFunc("/books", lt.ServeHTTP).Methods(http.MethodGet)

	ru := &handler.RegisterUser{
		Service:   &service.RegisterUser{DB: db, Repo: &rep},
		Validator: v,
	}
	r.HandleFunc("/register", ru.ServeHTTP).Methods(http.MethodPost)

	l := &handler.Login{
		Service: &service.Login{
			DB:             db,
			Repo:           &rep,
			TokenGenerator: jwter,
		},
		Validator: v,
	}
	r.HandleFunc("/login", l.ServeHTTP).Methods(http.MethodPost)

	return r, cleanup, nil
}

package main

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/naoyakurokawa/book_line_api/clock"
	"github.com/naoyakurokawa/book_line_api/config"
	"github.com/naoyakurokawa/book_line_api/handler"
	"github.com/naoyakurokawa/book_line_api/service"
	"github.com/naoyakurokawa/book_line_api/store"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	r := mux.NewRouter()

	clocker := clock.RealClocker{}
	db, cleanup, err := store.New(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}
	rep := store.Repository{Clocker: clocker}
	lt := &handler.ListBook{
		Service: &service.ListBook{DB: db, Repo: &rep},
	}
	r.HandleFunc("/books", lt.ServeHTTP).Methods(http.MethodGet)

	return r, cleanup, nil
}

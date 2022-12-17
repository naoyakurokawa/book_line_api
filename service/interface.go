package service

import (
	"context"

	"github.com/naoyakurokawa/book_line_api/entity"
	"github.com/naoyakurokawa/book_line_api/store"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . TaskAdder TaskLister UserRegister UserGetter TokenGenerator
type BookLister interface {
	ListBooks(ctx context.Context, db store.Queryer) (entity.Books, error)
}

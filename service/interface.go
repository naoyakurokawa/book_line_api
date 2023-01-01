package service

import (
	"context"

	"github.com/naoyakurokawa/book_line_api/entity"
	"github.com/naoyakurokawa/book_line_api/store"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . BookLister UserRegister
type BookLister interface {
	ListBooks(ctx context.Context, db store.Queryer) (entity.Books, error)
}

type UserRegister interface {
	RegisterUser(ctx context.Context, db store.Execer, u *entity.User) error
}

type UserGetter interface {
	GetUser(ctx context.Context, db store.Queryer, name string) (*entity.User, error)
}

type TokenGenerator interface {
	GenerateToken(ctx context.Context, u entity.User) ([]byte, error)
}

package service

import (
	"context"
	"fmt"

	"github.com/naoyakurokawa/book_line_api/entity"
	"github.com/naoyakurokawa/book_line_api/store"
)

type CreateBookService struct {
	DB   store.Execer
	Repo BookCreator
}

func (cb *CreateBookService) CreateBook(ctx context.Context, isbn int64) error {
	b := &entity.Book{
		Isbn: isbn,
	}
	if err := cb.Repo.CreateBook(ctx, cb.DB, b); err != nil {
		return fmt.Errorf("failed to create book: %w", err)
	}
	return nil
}

type FetchBooksService struct {
	DB   store.Queryer
	Repo BookFetcher
}

func (f *FetchBooksService) FetchBooks(ctx context.Context) (entity.Books, error) {
	bs, err := f.Repo.FetchBooks(ctx, f.DB)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}
	return bs, nil
}

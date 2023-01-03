package service

import (
	"context"
	"fmt"

	"github.com/naoyakurokawa/book_line_api/entity"
	"github.com/naoyakurokawa/book_line_api/store"
)

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

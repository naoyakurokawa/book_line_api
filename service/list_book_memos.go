package service

import (
	"context"
	"fmt"

	"github.com/naoyakurokawa/book_line_api/entity"
	"github.com/naoyakurokawa/book_line_api/store"
)

type ListBookMemos struct {
	DB   store.Queryer
	Repo BookMemoLister
}

func (lbm *ListBookMemos) ListBookMemos(ctx context.Context, bookID string) (entity.BookMemos, error) {
	bms, err := lbm.Repo.ListBookMemos(ctx, lbm.DB, bookID)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}
	return bms, nil
}

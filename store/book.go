package store

import (
	"context"

	"github.com/naoyakurokawa/book_line_api/entity"
)

func (r *Repository) ListBooks(
	ctx context.Context, db Queryer,
) (entity.Books, error) {
	books := entity.Books{}
	sql := `SELECT 
				id, isbn,
				created, modified 
			FROM books;`
	if err := db.SelectContext(ctx, &books, sql); err != nil {
		return nil, err
	}
	return books, nil
}

package store

import (
	"context"

	"github.com/naoyakurokawa/book_line_api/entity"
)

func (r *Repository) ListBookMemos(
	ctx context.Context, db Queryer, bookID string,
) (entity.BookMemos, error) {
	book_memos := entity.BookMemos{}
	sql := `SELECT 
				id, book_id, page, detail,
				created, modified 
			FROM book_memos WHERE book_id = ?`
	if err := db.SelectContext(ctx, &book_memos, sql, bookID); err != nil {
		return nil, err
	}
	return book_memos, nil
}

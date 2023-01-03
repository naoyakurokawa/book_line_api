package store

import (
	"context"

	"github.com/naoyakurokawa/book_line_api/entity"
)

func (r *Repository) FetchBooks(
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

func (r *Repository) CreateBook(
	ctx context.Context, db Execer, b *entity.Book,
) error {
	b.Created = r.Clocker.Now()
	b.Modified = r.Clocker.Now()
	sql := `INSERT INTO books (isbn, created, modified) 
		VALUES (?, ?, ?);`
	_, err := db.ExecContext(ctx, sql, b.Isbn, b.Created, b.Created)
	if err != nil {
		return err
	}
	return nil
}

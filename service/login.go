package service

import (
	"context"
	"fmt"

	"github.com/naoyakurokawa/book_line_api/store"
)

type Login struct {
	DB             store.Queryer
	Repo           UserFetcher
	TokenGenerator TokenGenerator
}

func (l *Login) Login(ctx context.Context, name, pw string) (string, error) {
	u, err := l.Repo.FetchUserByID(ctx, l.DB, name)
	if err != nil {
		return "", fmt.Errorf("failed to list: %w", err)
	}
	if err := u.ComparePassword(pw); err != nil {
		return "", fmt.Errorf("wrong password: %w", err)
	}
	jwt, err := l.TokenGenerator.GenerateToken(ctx, *u)
	if err != nil {
		return "", fmt.Errorf("failed to generate JWT: %w", err)
	}
	return string(jwt), nil
}

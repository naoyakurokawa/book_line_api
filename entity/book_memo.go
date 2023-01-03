package entity

import "time"

type BookMemoID int64

type BookMemo struct {
	ID       BookMemoID `json:"id" db:"id"`
	BookID   BookID     `json:"book_id" db:"book_id"`
	Page     int64      `json:"page" db:"page"`
	Detail   string     `json:"detail" db:"detail"`
	Created  time.Time  `json:"created" db:"created"`
	Modified time.Time  `json:"modified" db:"modified"`
}

type BookMemos []*BookMemo

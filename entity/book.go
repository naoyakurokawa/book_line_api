package entity

import "time"

type BookID int64
type Isbn int64

type Book struct {
	ID       BookID    `json:"id" db:"id"`
	Isbn     Isbn      `json:"isbn" db:"isbn"`
	Created  time.Time `json:"created" db:"created"`
	Modified time.Time `json:"modified" db:"modified"`
}

type Books []*Book

package entity

import "time"

type BookID int64

type Book struct {
	ID       BookID    `json:"id" db:"id"`
	Isbn     int64     `json:"isbn" db:"isbn"`
	Created  time.Time `json:"created" db:"created"`
	Modified time.Time `json:"modified" db:"modified"`
}

type Books []*Book

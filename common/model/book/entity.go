package book

import "time"

type Book struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Release     time.Time `json:"release"`
}

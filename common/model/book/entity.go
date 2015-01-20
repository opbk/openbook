package book

import "time"

type Book struct {
	Id          int64 `orm:"pk"`
	Title       string
	Pages       int
	Language    string
	Description string
	Release     time.Time
	Created     time.Time

	SeriesId     int64
	PublisherId  int64
	AuthorsId    []int64
	CategoriesId []int64
}

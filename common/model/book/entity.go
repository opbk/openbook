package book

import "time"

type Book struct {
	Id            int64
	Title         string
	Pages         int
	Language      string
	Short         string
	Description   string
	ServiceReview string
	CriticsReview string
	Release       time.Time
	Created       time.Time

	SeriesId     int64
	PublisherId  int64
	AuthorsId    []int64
	CategoriesId []int64
}

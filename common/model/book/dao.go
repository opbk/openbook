package book

import (
	"database/sql"
	"fmt"
	"strings"

	logger "github.com/cihub/seelog"

	"github.com/opbk/openbook/common/arrays"
	"github.com/opbk/openbook/common/db"
	"github.com/opbk/openbook/common/model/book/price"
)

const (
	FIND       = "SELECT id, title, pages, language, short, description, service_review, critics_review, release, created, series_id, publisher_id, array_agg(DISTINCT author_id) as authors_id, array_agg(DISTINCT category_id) as categories_id FROM books LEFT JOIN book_categories as bc ON id = bc.book_id LEFT JOIN author_books as ab ON id = ab.book_id WHERE id = $1 GROUP BY id"
	LIST_BY_ID = "SELECT id, title, pages, language, short, description, service_review, critics_review, release, created, series_id, publisher_id, array_agg(DISTINCT author_id) as authors_id, array_agg(DISTINCT category_id) as categories_id FROM books LEFT JOIN book_categories as bc ON id = bc.book_id LEFT JOIN author_books as ab ON id = ab.book_id WHERE id IN (%s) GROUP BY id"
	INSERT     = "INSERT INTO books(title, pages, language, short, description, service_review, critics_review, release, created, series_id, publisher_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id"
	UPDATE     = "UPDATE books SET title = $1, pages = $2, language = $3, language = $4, description = $5, service_review = $6, critics_review = $7, release = $8, release = $9, series_id = $10, publisher_id = $11 WHERE id = $12"
	DELETE     = "DELETE FROM books WHERE id = $1"
)

const (
	SEARCH       = "SELECT id, title, pages, language, short, description, service_review, critics_review, release, created, series_id, publisher_id, array_agg(DISTINCT author_id) as authors_id, array_agg(DISTINCT category_id) as categories_id FROM books LEFT JOIN book_categories as bc ON id = bc.book_id LEFT JOIN author_books as ab ON id = ab.book_id WHERE %s 1 = 1 GROUP BY id  ORDER BY created DESC LIMIT $1 OFFSET $2"
	SEARCH_COUNT = "SELECT COUNT(DISTINCT id) FROM books LEFT JOIN book_categories as bc ON id = bc.book_id LEFT JOIN author_books as ab ON id = ab.book_id WHERE %s 1 = 1"
)

const (
	ADD_BOOK_TO_CATEGORY = "INSERT INTO book_categories (book_id, category_id) VALUES ($1, $2)"
	ADD_BOOK_TO_AUTHOR   = "INSERT INTO author_books (book_id, author_id) VALUES ($1, $2)"
)

var searchWhere = map[string]string{
	"category":  " category_id = %d and ",
	"author":    " author_id = %d and ",
	"release":   " release > '%s' and ",
	"series":    " series_id = %d and ",
	"publisher": " publisher_id = %d and ",
	"search":    " title ILIKE '%%%s%%' and ",
}

func connection() *sql.DB {
	return db.Connection()
}

func scanRow(scanner db.RowScanner) *Book {
	var book *Book = new(Book)
	var categories string
	var authors string
	err := scanner.Scan(&book.Id, &book.Title, &book.Pages, &book.Language, &book.Short, &book.Description, &book.ServiceReview, &book.CriticsReview, &book.Release, &book.Created, &book.SeriesId, &book.PublisherId, &authors, &categories)
	if err != nil {
		logger.Errorf("Can't scan row: %s", err)
	}

	book.CategoriesId = db.StringToArray(categories)
	book.AuthorsId = db.StringToArray(authors)
	return book
}

func interateRows(rows *sql.Rows) []*Book {
	books := make([]*Book, 0)
	for rows.Next() {
		books = append(books, scanRow(rows))
	}

	return books
}

func iterateRowsToMap(rows *sql.Rows) map[int64]*Book {
	books := make(map[int64]*Book)
	for rows.Next() {
		book := scanRow(rows)
		books[book.Id] = book
	}

	return books
}

func MapById(ids []int64) map[int64]*Book {
	rows, err := connection().Query(fmt.Sprintf(LIST_BY_ID, strings.Join(arrays.Int64ToString(ids), ",")))
	if err != nil {
		logger.Errorf("Database error while getting map of books: %s", err)
	}

	return iterateRowsToMap(rows)
}

func Search(filter map[string]interface{}, limit, offset int) []*Book {
	var where string
	for key, val := range filter {
		where += fmt.Sprintf(searchWhere[key], val)
	}

	rows, err := connection().Query(fmt.Sprintf(SEARCH, where), limit, offset)
	if err != nil {
		logger.Errorf("Database error while searching list of books: %s", err)
	}

	return interateRows(rows)
}

func SearchCount(filter map[string]interface{}) int {
	var where string
	for key, val := range filter {
		where += fmt.Sprintf(searchWhere[key], val)
	}

	var count int
	row := connection().QueryRow(fmt.Sprintf(SEARCH_COUNT, where))
	err := row.Scan(&count)
	if err != nil {
		logger.Errorf("Database error while getting count of books: %s", err)
	}

	return count
}

func Find(id int64) *Book {
	row := connection().QueryRow(FIND, id)
	return scanRow(row)
}

func AddBookToCategory(bookId, categoryId int64) {
	_, err := connection().Exec(ADD_BOOK_TO_CATEGORY, bookId, categoryId)
	if err != nil {
		logger.Errorf("Database error while adding book %d to category %d: %s", bookId, categoryId, err)
	}
}

func AddBookToAuthor(bookId, authorId int64) {
	_, err := connection().Exec(ADD_BOOK_TO_AUTHOR, bookId, authorId)
	if err != nil {
		logger.Errorf("Database error while adding book %d to author %d: %s", bookId, authorId, err)
	}
}

func (b *Book) Prices() map[string]*price.Price {
	return price.MapByBookId(b.Id)
}

func (b *Book) Save() {
	if b.Id != 0 {
		b.update()
	} else {
		b.insert()
	}
}

func (b *Book) update() {
	_, err := connection().Exec(UPDATE, b.Title, b.Pages, b.Language, b.Short, b.Description, b.ServiceReview, b.CriticsReview, b.Release, b.Created, b.SeriesId, b.PublisherId, b.Id)
	if err != nil {
		logger.Errorf("Database error while updating book %d: %s", b.Id, err)
	}
}

func (b *Book) insert() {
	err := connection().QueryRow(INSERT, b.Title, b.Pages, b.Language, b.Short, b.Description, b.ServiceReview, b.CriticsReview, b.Release, b.Created, b.SeriesId, b.PublisherId).Scan(&b.Id)
	if err != nil {
		logger.Errorf("Database error while inserting book: %s", err)
	}
}

func (b *Book) Delete() {
	_, err := connection().Exec(DELETE, b.Id)
	if err != nil {
		logger.Errorf("Database error while deleting book %d: %s", b.Id, err)
	}
}

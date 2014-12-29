package book

import (
	"database/sql"
	"fmt"
	"strings"

	logger "github.com/cihub/seelog"
	"github.com/opbk/openbook/common/arrays"
	"github.com/opbk/openbook/common/db"
)

const (
	FIND                 = "SELECT id, title, description, release, created FROM books WHERE id = $1"
	LIST                 = "SELECT id, title, description, release, created FROM books LIMIT $1 OFFSET $2"
	INSERT               = "INSERT INTO books(title, description, release, created) VALUES ($1, $2, $3, $4) RETURNING id"
	UPDATE               = "UPDATE books SET title = $1, description = $2, release = $3, release = $4 WHERE id = $5"
	DELETE               = "DELETE FROM books WHERE id = $1"
	ADD_BOOK_TO_CATEGORY = "INSERT INTO book_categories (book_id, category_id) VALUES ($1, $2)"
	ADD_BOOK_TO_AUTHOR   = "INSERT INTO author_books (book_id, author_id) VALUES ($1, $2)"
)

const (
	SEARCH               = "SELECT id, title, description, release, created FROM books %s WHERE %s 1 = 1 ORDER BY created DESC LIMIT $1 OFFSET $2"
	SEARCH_JOIN_CATEGORY = " LEFT JOIN book_categories  as bc ON id = bc.book_id"
	SEARCH_JOIN_AUTHOR   = " LEFT JOIN author_books as ab ON id = ab.book_id"

	SEARCH_WHERE_CATEGORY = " category_id IN (%s) and "
	SEARCH_WHERE_AUTHOR   = " author_id = %d and "
	SEARCH_WHERE_RELEASE  = " release > '%s' and "
)

func connection() *sql.DB {
	return db.Connection()
}

func interateRows(rows *sql.Rows) []*Book {
	books := make([]*Book, 0)
	for rows.Next() {
		var book *Book = new(Book)
		rows.Scan(&book.Id, &book.Title, &book.Description, &book.Release, &book.Created)
		books = append(books, book)
	}
	return books
}

func List(limit, offset int) []*Book {
	rows, err := connection().Query(LIST, limit, offset)
	if err != nil {
		logger.Errorf("Database error while getting list of books: %s", err)
	}

	return interateRows(rows)
}

func Search(search map[string]interface{}, limit, offset int) []*Book {
	var join string
	var where string

	if _, ok := search["categories"]; ok {
		join += SEARCH_JOIN_CATEGORY
		where += fmt.Sprintf(SEARCH_WHERE_CATEGORY, strings.Join(arrays.Int64ToString(search["categories"].([]int64)), ","))
	}

	if _, ok := search["author"]; ok {
		join += SEARCH_JOIN_AUTHOR
		where += fmt.Sprintf(SEARCH_WHERE_AUTHOR, search["author"])
	}

	if _, ok := search["release"]; ok {
		where += fmt.Sprintf(SEARCH_WHERE_RELEASE, search["release"])
	}

	rows, err := connection().Query(fmt.Sprintf(SEARCH, join, where), limit, offset)
	if err != nil {
		logger.Errorf("Database error while getting list of books: %s", err)
	}

	return interateRows(rows)
}

func Find(id int64) *Book {
	var book *Book = new(Book)
	row := connection().QueryRow(FIND, id)
	err := row.Scan(&book.Id, &book.Title, &book.Description, &book.Release, &book.Created)

	if err != nil {
		logger.Errorf("Database error while finding book %d: %s", id, err)
		return nil
	}

	return book
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

func (b *Book) Save() {
	if b.Id != 0 {
		b.update()
	} else {
		b.insert()
	}
}

func (b *Book) update() {
	_, err := connection().Exec(UPDATE, b.Title, b.Description, b.Release, b.Created, b.Id)
	if err != nil {
		logger.Errorf("Database error while updating book %d: %s", b.Id, err)
	}
}

func (b *Book) insert() {
	err := connection().QueryRow(INSERT, b.Title, b.Description, b.Release, b.Created).Scan(&b.Id)
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

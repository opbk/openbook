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
	FIND                 = "SELECT id, title, description, release FROM books"
	LIST_BY_CATEGORY     = "SELECT id, title, description, release FROM books LEFT JOIN book_categories ON id = book_id WHERE category_id $1 LIMIT $2 OFFSET $3"
	LIST_BY_AUTHOR       = "SELECT id, title, description, release FROM books LEFT JOIN author_books ON id = book_id WHERE author_id = $1 LIMIT $2 OFFSET $3"
	INSERT               = "INSERT INTO books(title, description, release) VALUES ($1, $2, $3) RETURNING id"
	UPDATE               = "UPDATE books SET title = $1, description = $2, release = $3  WHERE id = $4"
	DELETE               = "DELETE FROM books WHERE id = $1"
	ADD_BOOK_TO_CATEGORY = "INSERT INTO book_categories (book_id, category_id) VALUES ($1, $2)"
	ADD_BOOK_TO_AUTHOR   = "INSERT INTO author_books (book_id, author_id) VALUES ($1, $2)"
)

func connection() *sql.DB {
	return db.Connection()
}

func interateRows(rows *sql.Rows) []*Book {
	books := make([]*Book, 0)
	for rows.Next() {
		var book *Book = new(Book)
		rows.Scan(&book.Id, &book.Title, &book.Description, &book.Release)
		books = append(books, book)
	}
	return books
}

func ListByCategory(categoryId int64, limit, offset int) []*Book {
	return ListByCategories([]int64{categoryId}, limit, offset)
}

func ListByCategories(categoriesId []int64, limit, offset int) []*Book {
	rows, err := connection().Query(fmt.Sprintf(LIST_BY_CATEGORY, strings.Join(arrays.Int64ToString(categoriesId), ",")), limit, offset)
	if err != nil {
		logger.Errorf("Database error while getting list of books categories: %s", err)
	}

	return interateRows(rows)
}

func ListByAuthor(authorId int64, limit, offset int) []*Book {
	return ListByCategories([]int64{authorId}, limit, offset)
}

func ListByAuthors(authorsId []int64, limit, offset int) []*Book {
	rows, err := connection().Query(fmt.Sprintf(LIST_BY_CATEGORY, strings.Join(arrays.Int64ToString(authorsId), ",")), limit, offset)
	if err != nil {
		logger.Errorf("Database error while getting list of books by authors: %s", err)
	}

	return interateRows(rows)
}

func Find(id int64) *Book {
	var book *Book = new(Book)
	row := connection().QueryRow(FIND, id)
	err := row.Scan(&book.Id, &book.Title, &book.Description, &book.Release)

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

func (b *Book) Save() {
	if b.Id != 0 {
		b.update()
	} else {
		b.insert()
	}
}

func (b *Book) update() {
	_, err := connection().Exec(UPDATE, b.Title, b.Description, b.Release, b.Id)
	if err != nil {
		logger.Errorf("Database error while updating book %d: %s", b.Id, err)
	}
}

func (b *Book) insert() {
	err := connection().QueryRow(INSERT, b.Title, b.Description, b.Release).Scan(&b.Id)
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

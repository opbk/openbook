package utils

import (
	"database/sql"

	logger "github.com/cihub/seelog"

	"github.com/opbk/openbook/common/db"
	"github.com/opbk/openbook/common/model/book"
	"github.com/opbk/openbook/common/model/book/category"
)

const (
	SELECT_BOOK_CATEGORIES = "SELECT category_id from book_categories WHERE book_id = $1"
	DELETE_CONNECTION      = "DELETE FROM book_categories WHERE book_id = $1 AND category_id = $2"
)

func connection() *sql.DB {
	return db.Connection()
}

func FixCategories(b *book.Book) {
	var categories = GetBookCategories(b)

	for _, c := range categories {
		for _, cat := range c.GetPath() {
			if notIn(cat.Id, categories) {
				book.AddBookToCategory(b.Id, cat.Id)
			}
		}
	}
}

func DeleteConnections(book *book.Book, category *category.Category) {
	var categories = category.GetPath()
	categories = append(categories, category)

	for _, c := range categories {
		_, err := connection().Query(DELETE_CONNECTION, book.Id, c.Id)
		if err != nil {
			logger.Errorf("Database error while delete connection %d to %d: %s", book.Id, c.Id, err)
		}
	}
}

func notIn(id int64, categories []*category.Category) bool {
	for _, c := range categories {
		if id == c.Id {
			return false
		}
	}
	return true
}

func GetBookCategories(b *book.Book) []*category.Category {
	var ids = make([]int64, 0)
	var categories = make([]*category.Category, 0)

	rows, err := connection().Query(SELECT_BOOK_CATEGORIES, b.Id)
	if err != nil {
		logger.Errorf("Database error while getting book categories %d: %s", b.Id, err)
	}

	for rows.Next() {
		var id int64
		rows.Scan(&id)
		ids = append(ids, id)
	}

	for _, i := range ids {
		c := category.Find(i)
		categories = append(categories, c)
	}

	return categories
}

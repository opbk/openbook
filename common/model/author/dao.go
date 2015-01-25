package author

import (
	"database/sql"
	"fmt"
	"strings"

	logger "github.com/cihub/seelog"
	"github.com/opbk/openbook/common/arrays"
	"github.com/opbk/openbook/common/db"
)

const (
	LIST       = "SELECT id, name, description, books FROM authors"
	LIST_BY_ID = "SELECT id, name, description, books FROM authors WHERE id IN (%s)"
	FIND       = "SELECT id, name, description, books FROM authors WHERE id = $1"
	INSERT     = "INSERT INTO authors (name, description, books) VALUES ($1, $2, $3) RETURNING id"
	UPDATE     = "UPDATE authors SET name = $1, description = $2, books = $3 WHERE id = $4"
	DELETE     = "DELETE FROM authors WHERE id = $1"
)

const (
	SEARCH = "SELECT a.id, a.name, a.description, COUNT(DISTINCT b.id) as books FROM authors as a LEFT JOIN author_books as ab ON a.id = ab.author_id LEFT JOIN books as b ON b.id = ab.book_id LEFT JOIN book_categories as bc ON b.id = bc.book_id WHERE %s 1 = 1 GROUP BY a.id  ORDER BY books DESC"
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

func scanRow(scanner db.RowScanner) *Author {
	var author *Author = new(Author)
	err := scanner.Scan(&author.Id, &author.Name, &author.Description, &author.Books)
	if err != nil {
		logger.Errorf("Can't scan row: %s", err)
	}

	return author
}

func interateRows(rows *sql.Rows) []*Author {
	authors := make([]*Author, 0)
	for rows.Next() {
		authors = append(authors, scanRow(rows))
	}

	return authors
}

func interateRowsToMap(rows *sql.Rows) map[int64]*Author {
	authors := make(map[int64]*Author)
	for rows.Next() {
		author := scanRow(rows)
		authors[author.Id] = author
	}

	return authors
}

func List() []*Author {
	rows, err := connection().Query(LIST)
	if err != nil {
		logger.Errorf("Database error while getting list of authors: %s", err)
	}

	return interateRows(rows)
}

func MapById(ids []int64) map[int64]*Author {
	rows, err := connection().Query(fmt.Sprintf(LIST_BY_ID, strings.Join(arrays.Int64ToString(ids), ",")))
	if err != nil {
		logger.Errorf("Database error while getting list of authors by id: %s", err)
	}

	return interateRowsToMap(rows)
}

func Search(filter map[string]interface{}) []*Author {
	var where string
	for key, val := range filter {
		where += fmt.Sprintf(searchWhere[key], val)
	}

	rows, err := connection().Query(fmt.Sprintf(SEARCH, where))
	if err != nil {
		logger.Errorf("Database error while searching list of authors: %s", err)
	}

	return interateRows(rows)
}

func Find(id int64) *Author {
	row := connection().QueryRow(FIND, id)
	author := scanRow(row)

	return author
}

func (a *Author) Save() {
	if a.Id != 0 {
		a.update()
	} else {
		a.insert()
	}
}

func (a *Author) update() {
	_, err := connection().Exec(UPDATE, a.Name, a.Description, a.Books, a.Id)
	if err != nil {
		logger.Errorf("Database error while updating author %d: %s", a.Id, err)
	}
}

func (a *Author) insert() {
	err := connection().QueryRow(INSERT, a.Name, a.Description, a.Books).Scan(&a.Id)
	if err != nil {
		logger.Errorf("Database error while inserting author: %s", err)
	}
}

func (a *Author) Delete() {
	_, err := connection().Exec(DELETE, a.Id)
	if err != nil {
		logger.Errorf("Database error while deleting author %d: %s", a.Id, err)
	}
}

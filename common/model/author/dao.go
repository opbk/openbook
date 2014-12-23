package author

import (
	"database/sql"
	logger "github.com/cihub/seelog"
	"github.com/opbk/openbook/common/db"
)

const (
	LIST       = "SELECT id, name FROM authors"
	FIND_BY_ID = "SELECT id, name FROM authors WHERE id = $1"
	INSERT     = "INSERT INTO authors (name) VALUES ($1) RETURNING id"
	UPDATE     = "UPDATE authors SET name = $1 WHERE id = $2"
	DELETE     = "DELETE FROM authors WHERE id = $1"
)

func connection() *sql.DB {
	return db.Connection()
}

func interateRows(rows *sql.Rows) []*Author {
	authors := make([]*Author, 0)
	for rows.Next() {
		var author *Author = new(Author)
		rows.Scan(&author.Id, &author.Name)
		authors = append(authors, author)
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

func Find(id int64) *Author {
	var author *Author = new(Author)
	row := connection().QueryRow(FIND_BY_ID, id)
	err := row.Scan(&author.Id, &author.Name)

	if err != nil {
		logger.Errorf("Database error while finding author %d: %s", id, err)
		return nil
	}

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
	_, err := connection().Exec(UPDATE, a.Name, a.Id)
	if err != nil {
		logger.Errorf("Database error while updating author %d: %s", a.Id, err)
	}
}

func (a *Author) insert() {
	err := connection().QueryRow(INSERT, a.Name).Scan(&a.Id)
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

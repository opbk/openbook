package publisher

import (
	"database/sql"
	"fmt"
	"strings"

	logger "github.com/cihub/seelog"
	"github.com/opbk/openbook/common/arrays"
	"github.com/opbk/openbook/common/db"
)

const (
	LIST       = "SELECT id, name, description, books FROM publishers"
	LIST_BY_ID = "SELECT id, name, description, books FROM publishers WHERE id IN (%s)"
	FIND       = "SELECT id, name, description, books FROM publishers WHERE id = $1"
	INSERT     = "INSERT INTO publishers (name, description, books) VALUES ($1, $2, $3) RETURNING id"
	UPDATE     = "UPDATE publishers SET name = $1, description = $2, books = $3 WHERE id = $4"
	DELETE     = "DELETE FROM publishers WHERE id = $1"
)

const (
	SEARCH = "SELECT p.id, p.name, p.description, COUNT(DISTINCT b.id) as books FROM publishers as p LEFT JOIN books as b ON p.id = b.publisher_id LEFT JOIN book_categories as bc ON b.id = bc.book_id LEFT JOIN author_books as ab ON b.id = ab.book_id WHERE %s 1 = 1 GROUP BY p.id ORDER BY books DESC"
)

var searchWhere = map[string]string{
	"category":  " category_id = %d and ",
	"author":    " author_id = %d and ",
	"release":   " release > '%s' and ",
	"series":    " series_id = %d and ",
	"publisher": " publisher_id = %d and ",
	"search":    " title LIKE '%%%s%%' and ",
}

func connection() *sql.DB {
	return db.Connection()
}

func interateRows(rows *sql.Rows) []*Publisher {
	publishers := make([]*Publisher, 0)
	for rows.Next() {
		var publisher *Publisher = new(Publisher)
		rows.Scan(&publisher.Id, &publisher.Name, &publisher.Description, &publisher.Books)
		publishers = append(publishers, publisher)
	}

	return publishers
}

func interateRowsToMap(rows *sql.Rows) map[int64]*Publisher {
	publishers := make(map[int64]*Publisher)
	for rows.Next() {
		var publisher *Publisher = new(Publisher)
		rows.Scan(&publisher.Id, &publisher.Name, &publisher.Description, &publisher.Books)
		publishers[publisher.Id] = publisher
	}

	return publishers
}

func List() []*Publisher {
	rows, err := connection().Query(LIST)
	if err != nil {
		logger.Errorf("Database error while getting list of publishers: %s", err)
	}

	return interateRows(rows)
}

func MapById(ids []int64) map[int64]*Publisher {
	rows, err := connection().Query(fmt.Sprintf(LIST_BY_ID, strings.Join(arrays.Int64ToString(ids), ",")))
	if err != nil {
		logger.Errorf("Database error while getting list of publishers by id: %s", err)
	}

	return interateRowsToMap(rows)
}

func Search(search map[string]interface{}) []*Publisher {
	var where string
	for key, val := range search {
		where += fmt.Sprintf(searchWhere[key], val)
	}

	rows, err := connection().Query(fmt.Sprintf(SEARCH, where))
	if err != nil {
		logger.Errorf("Database error while searching list of books: %s", err)
	}

	return interateRows(rows)
}

func Find(id int64) *Publisher {
	var publisher *Publisher = new(Publisher)
	row := connection().QueryRow(FIND, id)
	err := row.Scan(&publisher.Id, &publisher.Name, &publisher.Description, &publisher.Books)

	if err != nil {
		logger.Errorf("Database error while finding publisher %d: %s", id, err)
		return nil
	}

	return publisher
}

func (p *Publisher) Save() {
	if p.Id != 0 {
		p.update()
	} else {
		p.insert()
	}
}

func (p *Publisher) update() {
	_, err := connection().Exec(UPDATE, p.Name, p.Description, p.Books, p.Id)
	if err != nil {
		logger.Errorf("Database error while updating publisher %d: %s", p.Id, err)
	}
}

func (p *Publisher) insert() {
	err := connection().QueryRow(INSERT, p.Name, p.Description, p.Books).Scan(&p.Id)
	if err != nil {
		logger.Errorf("Database error while inserting publisher: %s", err)
	}
}

func (p *Publisher) Delete() {
	_, err := connection().Exec(DELETE, p.Id)
	if err != nil {
		logger.Errorf("Database error while deleting publisher %d: %s", p.Id, err)
	}
}

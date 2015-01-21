package pricetype

import (
	"database/sql"

	logger "github.com/cihub/seelog"
	"github.com/opbk/openbook/common/db"
)

const (
	LIST   = "SELECT id, type, name FROM prices"
	FIND   = "SELECT id, type, name FROM prices WHERE id = $1"
	INSERT = "INSERT INTO prices (type, name) VALUES ($1, $2) RETURNING id"
	UPDATE = "UPDATE prices SET type = $1, name = $2 WHERE id = $3"
	DELETE = "DELETE FROM prices WHERE id = $1"
)

func connection() *sql.DB {
	return db.Connection()
}

func scanRow(scanner db.RowScanner) *Type {
	var priceType *Type = new(Type)
	err := scanner.Scan(&priceType.Id, &priceType.Type, &priceType.Name)
	if err != nil {
		logger.Errorf("Can't scan row: %s", err)
	}

	return priceType
}

func interateRows(rows *sql.Rows) []*Type {
	priceTypes := make([]*Type, 0)
	for rows.Next() {
		priceTypes = append(priceTypes, scanRow(rows))
	}

	return priceTypes
}

func List() []*Type {
	rows, err := connection().Query(LIST)
	if err != nil {
		logger.Errorf("Database error while getting list of priceTypes: %s", err)
	}

	return interateRows(rows)
}

func Find(id int64) *Type {
	row := connection().QueryRow(FIND, id)
	priceType := scanRow(row)

	return priceType
}

func (a *Type) Save() {
	if a.Id != 0 {
		a.update()
	} else {
		a.insert()
	}
}

func (a *Type) update() {
	_, err := connection().Exec(UPDATE, a.Type, a.Name, a.Id)
	if err != nil {
		logger.Errorf("Database error while updating priceType %d: %s", a.Id, err)
	}
}

func (a *Type) insert() {
	err := connection().QueryRow(INSERT, a.Type, a.Name).Scan(&a.Id)
	if err != nil {
		logger.Errorf("Database error while inserting priceType: %s", err)
	}
}

func (a *Type) Delete() {
	_, err := connection().Exec(DELETE, a.Id)
	if err != nil {
		logger.Errorf("Database error while deleting priceType %d: %s", a.Id, err)
	}
}

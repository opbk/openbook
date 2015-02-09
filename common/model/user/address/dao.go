package address

import (
	"database/sql"

	logger "github.com/cihub/seelog"

	"github.com/opbk/openbook/common/db"
)

const (
	LIST_BY_USER = "SELECT id, user_id, address, comment FROM addresses WHERE user_id = $1"
	FIND         = "SELECT id, user_id, address, comment FROM addresses WHERE id = $1"
	INSERT       = "INSERT INTO addresses (user_id, address, comment) VALUES ($1, $2, $3) RETURNING id"
	UPDATE       = "UPDATE addresses SET user_id = $1, address = $2, comment = $3 WHERE id = $4"
	DELETE       = "DELETE FROM addresses WHERE id = $1"
)

func connection() *sql.DB {
	return db.Connection()
}

func scanRow(scaner db.RowScanner) *Address {
	var address *Address = new(Address)
	err := scaner.Scan(&address.Id, &address.UserId, &address.Address, &address.Comment)
	if err != nil {
		logger.Errorf("Can't scan row: %s", err)
	}

	return address
}

func interateRows(rows *sql.Rows) []*Address {
	addresses := make([]*Address, 0)
	for rows.Next() {
		addresses = append(addresses, scanRow(rows))
	}

	return addresses
}

func interateRowsToMap(rows *sql.Rows) map[int64]*Address {
	addresses := make(map[int64]*Address)
	for rows.Next() {
		address := scanRow(rows)
		addresses[address.Id] = address
	}

	return addresses
}

func ListByUser(id int64) []*Address {
	rows, err := connection().Query(LIST_BY_USER, id)
	if err != nil {
		logger.Errorf("Database error while getting list of addresses by user: %s", err)
	}

	return interateRows(rows)
}

func MapByUser(id int64) map[int64]*Address {
	rows, err := connection().Query(LIST_BY_USER, id)
	if err != nil {
		logger.Errorf("Database error while getting list of addresses by user: %s", err)
	}

	return interateRowsToMap(rows)
}

func Find(id int64) *Address {
	return scanRow(connection().QueryRow(FIND, id))
}

func (a *Address) Save() {
	if a.Id != 0 {
		a.update()
	} else {
		a.insert()
	}
}

func (a *Address) update() {
	_, err := connection().Exec(UPDATE, a.UserId, a.Address, a.Comment, a.Id)
	if err != nil {
		logger.Errorf("Database error while updating address %d: %s", a.Id, err)
	}
}

func (a *Address) insert() {
	err := connection().QueryRow(INSERT, a.UserId, a.Address, a.Comment).Scan(&a.Id)
	if err != nil {
		logger.Errorf("Database error while inserting address: %s", err)
	}
}

func (a *Address) Delete() {
	_, err := connection().Exec(DELETE, a.Id)
	if err != nil {
		logger.Errorf("Database error while deleting address %d: %s", a.Id, err)
	}
}

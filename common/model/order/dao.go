package order

import (
	"database/sql"

	logger "github.com/cihub/seelog"

	"github.com/opbk/openbook/common/db"
)

const (
	LIST_BY_USER = "SELECT id, user_id, address_id, status, comment, created, modified FROM orders WHERE user_id = $1"
	FIND         = "SELECT id, user_id, address_id, status, comment, created, modified FROM orders WHERE id = $1"
	INSERT       = "INSERT INTO orders (user_id, address_id, status, comment) VALUES ($1, $2, $3, $4) RETURNING id"
	UPDATE       = "UPDATE orders SET user_id = $1, address_id = $2, status = $3, comment = $4 WHERE id = $5"
	DELETE       = "DELETE FROM orders WHERE id = $1"
)

func connection() *sql.DB {
	return db.Connection()
}

func scanRow(scaner db.RowScanner) *Order {
	var order *Order = new(Order)
	err := scaner.Scan(&order.Id, &order.UserId, &order.AddressId, &order.Status, &order.Comment, &order.Created, &order.Modified)
	if err != nil {
		logger.Errorf("Can't scan row: %s", err)
	}

	return order
}

func interateRows(rows *sql.Rows) []*Order {
	orders := make([]*Order, 0)
	for rows.Next() {
		orders = append(orders, scanRow(rows))
	}

	return orders
}

func interateRowsToMap(rows *sql.Rows) map[int64]*Order {
	orders := make(map[int64]*Order)
	for rows.Next() {
		order := scanRow(rows)
		orders[order.Id] = order
	}

	return orders
}

func ListByUser(id int64) []*Order {
	rows, err := connection().Query(LIST_BY_USER, id)
	if err != nil {
		logger.Errorf("Database error while getting list of orders by user: %s", err)
	}

	return interateRows(rows)
}

func Find(id int64) *Order {
	return scanRow(connection().QueryRow(FIND, id))
}

func (o *Order) Save() {
	if o.Id != 0 {
		o.update()
	} else {
		o.insert()
	}
}

func (o *Order) update() {
	_, err := connection().Exec(UPDATE, o.UserId, o.AddressId, o.Status, o.Comment, o.Id)
	if err != nil {
		logger.Errorf("Database error while updating order %d: %s", o.Id, err)
	}
}

func (o *Order) insert() {
	err := connection().QueryRow(INSERT, o.UserId, o.AddressId, o.Status, o.Comment).Scan(&o.Id)
	if err != nil {
		logger.Errorf("Database error while inserting order: %s", err)
	}
}

func (o *Order) Delete() {
	_, err := connection().Exec(DELETE, o.Id)
	if err != nil {
		logger.Errorf("Database error while deleting order %d: %s", o.Id, err)
	}
}

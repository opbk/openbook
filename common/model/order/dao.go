package order

import (
	"database/sql"

	logger "github.com/cihub/seelog"

	"github.com/opbk/openbook/common/db"
)

const (
	FIND           = "SELECT id, book_id, user_id, address_id, status, comment, created, modified FROM orders LEFT JOIN book_orders ON order_id = id WHERE id = $1"
	INSERT         = "INSERT INTO orders (user_id, address_id, status, comment) VALUES ($1, $2, $3, $4) RETURNING id"
	UPDATE         = "UPDATE orders SET user_id = $1, address_id = $2, status = $3, comment = $4 WHERE id = $5"
	DELETE         = "DELETE FROM orders WHERE id = $1"
	ADD_ORDER_BOOK = "INSERT INTO book_orders (order_id, book_id) VALUES ($1, $2)"

	LIST_BY_USER  = "SELECT id, book_id, user_id, address_id, status, comment, created, modified FROM orders LEFT JOIN book_orders ON order_id = id WHERE user_id = $1 ORDER BY created DESC LIMIT $2 OFFSET $3"
	COUNT_BY_USER = "SELECT COUNT(*) FROM orders WHERE user_id = $1"

	LIST_BY_USER_STATUS  = "SELECT id, book_id, user_id, address_id, status, comment, created, modified FROM orders LEFT JOIN book_orders ON order_id = id WHERE user_id = $1 and status = $2 ORDER BY created DESC LIMIT $3 OFFSET $4"
	COUNT_BY_USER_STATUS = "SELECT COUNT(*) FROM orders WHERE user_id = $1 and status = $2"
)

func connection() *sql.DB {
	return db.Connection()
}

func scanRow(scaner db.RowScanner) *Order {
	var order *Order = new(Order)
	err := scaner.Scan(&order.Id, &order.BookId, &order.UserId, &order.AddressId, &order.Status, &order.Comment, &order.Created, &order.Modified)
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

func ListByUserAndStatusWithLimit(id int64, status string, limit, offset int) []*Order {
	var rows *sql.Rows
	var err error

	if status != "" {
		rows, err = connection().Query(LIST_BY_USER_STATUS, id, status, limit, offset)
	} else {
		rows, err = connection().Query(LIST_BY_USER, id, limit, offset)
	}

	if err != nil {
		logger.Errorf("Database error while getting list of orders by user: %s", err)
	}

	return interateRows(rows)
}

func CountByUserAndStatus(id int64, status string) int {
	var count int
	var err error

	if status != "" {
		err = connection().QueryRow(COUNT_BY_USER_STATUS, id, status).Scan(&count)
	} else {
		err = connection().QueryRow(COUNT_BY_USER, id).Scan(&count)
	}

	if err != nil {
		logger.Errorf("Database error while getting count of orders by user: %s", err)
	}

	return count
}

func AddOrderBook(id, bookId int64) {
	_, err := connection().Exec(ADD_ORDER_BOOK, id, bookId)
	if err != nil {
		logger.Errorf("Database error while inserting order: %s", err)
	}
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

func (o *Order) Return() {
	o.Status = RETURNING
	o.Save()
}

func (o *Order) Delete() {
	_, err := connection().Exec(DELETE, o.Id)
	if err != nil {
		logger.Errorf("Database error while deleting order %d: %s", o.Id, err)
	}
}

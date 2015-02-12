package order

import (
	"database/sql"

	logger "github.com/cihub/seelog"

	"github.com/opbk/openbook/common/db"
)

const (
	LIST_BY_USER   = "SELECT id, book_id, user_id, address_id, status, comment, created, modified FROM orders LEFT JOIN book_orders ON order_id = id WHERE user_id = $1 ORDER BY created DESC LIMIT $2 OFFSET $3"
	COUNT_BY_USER  = "SELECT COUNT(*) FROM orders WHERE user_id = $1"
	FIND           = "SELECT id, book_id, user_id, address_id, status, comment, created, modified FROM orders LEFT JOIN book_orders ON order_id = id WHERE id = $1"
	INSERT         = "INSERT INTO orders (user_id, address_id, status, comment) VALUES ($1, $2, $3, $4) RETURNING id"
	UPDATE         = "UPDATE orders SET user_id = $1, address_id = $2, status = $3, comment = $4 WHERE id = $5"
	DELETE         = "DELETE FROM orders WHERE id = $1"
	ADD_ORDER_BOOK = "INSERT INTO book_orders (order_id, book_id) VALUES ($1, $2)"
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

func ListByUserWithLimit(id int64, limit, offset int) []*Order {
	rows, err := connection().Query(LIST_BY_USER, id, limit, offset)
	if err != nil {
		logger.Errorf("Database error while getting list of orders by user: %s", err)
	}

	return interateRows(rows)
}

func CountByUser(id int64) int {
	var count int
	err := connection().QueryRow(COUNT_BY_USER, id).Scan(&count)
	if err != nil {
		logger.Errorf("Database error while getting count of orders by user: %s", err)
	}

	return count
}

func AddOrderBook(id, bookId int) {
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

func (o *Order) Delete() {
	_, err := connection().Exec(DELETE, o.Id)
	if err != nil {
		logger.Errorf("Database error while deleting order %d: %s", o.Id, err)
	}
}

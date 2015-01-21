package price

import (
	"database/sql"

	logger "github.com/cihub/seelog"
	"github.com/opbk/openbook/common/db"
)

const (
	ADD_BOOK_PRICE = "INSERT INTO book_prices (book_id, price_type_id, price) VALUES ($1, $2, $3)"
	PRICE_LIST     = "SELECT p.id, p.type, p.name, bp.price FROM book_prices as bp LEFT JOIN prices as p ON p.id = bp.price_type_id WHERE bp.book_id = $1"
)

func connection() *sql.DB {
	return db.Connection()
}

func AddBookPrice(bookId, typeId int64, price float64) {
	_, err := connection().Exec(ADD_BOOK_PRICE, bookId, typeId, price)
	if err != nil {
		logger.Errorf("Database error while adding price %d to book %d: %s", typeId, bookId, err)
	}
}

func MapByBookId(id int64) map[string]*Price {
	rows, err := connection().Query(PRICE_LIST, id)
	if err != nil {
		logger.Errorf("Database error while searching list of books: %s", err)
	}

	prices := make(map[string]*Price)
	for rows.Next() {
		price := new(Price)
		rows.Scan(&price.Id, &price.Type.Type, &price.Name, &price.Price)
		prices[price.Type.Type] = price
	}

	return prices
}

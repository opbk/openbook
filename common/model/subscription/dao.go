package subscription

import (
	"database/sql"

	logger "github.com/cihub/seelog"

	"github.com/opbk/openbook/common/db"
)

const (
	LIST         = "SELECT id, name, description, price, enabled FROM subscriptions"
	LIST_ENABLED = "SELECT id, name, description, price, enabled FROM subscriptions WHERE enabled = true"
	FIND         = "SELECT id, name, description, price, enabled FROM subscriptions WHERE id = $1"
	INSERT       = "INSERT INTO subscriptions (name, description, price, enabled) VALUES ($1, $2, $3, $4) RETURNING id"
	UPDATE       = "UPDATE subscriptions SET name = $1, description = $2, price = $3, enabled = $4  WHERE id = $5"
	DELETE       = "DELETE FROM subscriptions WHERE id = $1"
)

func connection() *sql.DB {
	return db.Connection()
}

func scanRow(scaner db.RowScanner) *Subscription {
	var subscription *Subscription = new(Subscription)
	err := scaner.Scan(&subscription.Id, &subscription.Name, &subscription.Description, &subscription.Price, &subscription.Enabled)
	if err != nil {
		logger.Errorf("Can't scan row: %s", err)
		return nil
	}

	return subscription
}

func interateRows(rows *sql.Rows) []*Subscription {
	users := make([]*Subscription, 0)
	for rows.Next() {
		users = append(users, scanRow(rows))
	}

	return users
}

func List() []*Subscription {
	rows, err := connection().Query(LIST)
	if err != nil {
		logger.Errorf("Database error while getting list of users: %s", err)
	}

	return interateRows(rows)
}

func ListEnabled() []*Subscription {
	rows, err := connection().Query(LIST_ENABLED)
	if err != nil {
		logger.Errorf("Database error while getting list of users: %s", err)
	}

	return interateRows(rows)
}

func Find(id int64) *Subscription {
	return scanRow(connection().QueryRow(FIND, id))
}

func (s *Subscription) Save() {
	if s.Id != 0 {
		s.update()
	} else {
		s.insert()
	}
}

func (s *Subscription) update() {
	_, err := connection().Exec(UPDATE, s.Name, s.Description, s.Price, s.Enabled, s.Id)
	if err != nil {
		logger.Errorf("Database error while updating subscription %d: %s", s.Id, err)
	}
}

func (s *Subscription) insert() {
	err := connection().QueryRow(INSERT, s.Name, s.Description, s.Price, s.Enabled).Scan(&s.Id)
	if err != nil {
		logger.Errorf("Database error while inserting subscription: %s", err)
	}
}

func (s *Subscription) Delete() {
	_, err := connection().Exec(DELETE, s.Id)
	if err != nil {
		logger.Errorf("Database error while deleting subscription %d: %s", s.Id, err)
	}
}

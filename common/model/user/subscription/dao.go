package subscription

import (
	"database/sql"

	logger "github.com/cihub/seelog"

	"github.com/opbk/openbook/common/db"
)

const (
	LIST         = "SELECT id, name, description, price, user_id, expiration FROM subscriptions LEFT JOIN user_subscriptions ON id = subscription_id"
	FIND_BY_USER = "SELECT id, name, description, price, user_id, expiration FROM subscriptions LEFT JOIN user_subscriptions ON id = subscription_id WHERE user_id = $1"
	INSERT       = "INSERT INTO user_subscriptions (user_id, subscription_id, expiration) VALUES ($1, $2, $3)"
	DELETE       = "DELETE FROM user_subscriptions WHERE user_id = $1 and subscription_id = $2"
)

func connection() *sql.DB {
	return db.Connection()
}

func scanRow(scaner db.RowScanner) *UserSubscription {
	var us *UserSubscription = new(UserSubscription)
	err := scaner.Scan(&us.Id, &us.Name, &us.Description, &us.Price, &us.UserId, &us.Expiration)
	if err != nil {
		logger.Errorf("Can't scan row: %s", err)
		return nil
	}

	return us
}

func interateRows(rows *sql.Rows) []*UserSubscription {
	subscriptions := make([]*UserSubscription, 0)
	for rows.Next() {
		subscriptions = append(subscriptions, scanRow(rows))
	}

	return subscriptions
}

func List() []*UserSubscription {
	rows, err := connection().Query(LIST)
	if err != nil {
		logger.Errorf("Database error while getting list of users subscription: %s", err)
	}

	return interateRows(rows)
}

func FindByUser(id int64) *UserSubscription {
	return scanRow(connection().QueryRow(FIND_BY_USER, id))
}

func (us *UserSubscription) Insert() {
	_, err := connection().Exec(INSERT, us.UserId, us.Id, us.Expiration)
	if err != nil {
		logger.Errorf("Database error while inserting user subscription: %s", err)
	}
}

func (us *UserSubscription) Delete() {
	_, err := connection().Exec(DELETE, us.UserId, us.Id)
	if err != nil {
		logger.Errorf("Database error while deleting subscription %d: %s", us.Id, err)
	}
}

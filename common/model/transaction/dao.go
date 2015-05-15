package transaction

import (
	"database/sql"

	logger "github.com/cihub/seelog"

	"github.com/opbk/openbook/common/db"
)

const (
	FIND             = "SELECT id, user_id, subscription_id, payload, status FROM transactions WHERE id = $1"
	FIND_NEW_BY_USER = "SELECT id, user_id, subscription_id, payload, status FROM transactions WHERE user_id = $1 and status = 'new'"
	INSERT           = "INSERT INTO transactions (user_id, subscription_id, payload, status) VALUES ($1, $2, $3, $4) RETURNING id"
	UPDATE           = "UPDATE transactions SET user_id = $1, subscription_id = $2, payload = $3, status = $4 WHERE id = $5"
	DELETE           = "DELETE FROM transactions WHERE id = $1"
)

func connection() *sql.DB {
	return db.Connection()
}

func scanRow(scaner db.RowScanner) *Transaction {
	var transaction *Transaction = new(Transaction)
	err := scaner.Scan(&transaction.Id, &transaction.UserId, &transaction.SubscriptionId, &transaction.Payload, &transaction.Status)
	if err != nil {
		logger.Errorf("Can't scan row: %s", err)
		return nil
	}

	return transaction
}

func interateRows(rows *sql.Rows) []*Transaction {
	users := make([]*Transaction, 0)
	for rows.Next() {
		users = append(users, scanRow(rows))
	}

	return users
}

func Find(id int64) *Transaction {
	return scanRow(connection().QueryRow(FIND, id))
}

func NewTransaction(userId int64) *Transaction {
	t := scanRow(connection().QueryRow(FIND_NEW_BY_USER, userId))
	if t == nil {
		t = &Transaction{
			UserId: userId,
			Status: NEW,
		}
	}
	return t
}

func (t *Transaction) Save() {
	if t.Id != 0 {
		t.update()
	} else {
		t.insert()
	}
}

func (t *Transaction) update() {
	_, err := connection().Exec(UPDATE, t.UserId, t.SubscriptionId, t.Payload, t.Status, t.Id)
	if err != nil {
		logger.Errorf("Database error while updating subscription %d: %s", t.Id, err)
	}
}

func (t *Transaction) insert() {
	err := connection().QueryRow(INSERT, t.UserId, t.SubscriptionId, t.Payload, t.Status).Scan(&t.Id)
	if err != nil {
		logger.Errorf("Database error while inserting subscription: %s", err)
	}
}

func (t *Transaction) Delete() {
	_, err := connection().Exec(DELETE, t.Id)
	if err != nil {
		logger.Errorf("Database error while deleting subscription %d: %s", t.Id, err)
	}
}

func (t *Transaction) Complite() {
	t.Status = COMPLITED
	t.Save()
}

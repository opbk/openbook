package user

import (
	"database/sql"

	logger "github.com/cihub/seelog"

	"github.com/opbk/openbook/common/db"
	"github.com/opbk/openbook/common/model/user/subscription"
)

const (
	LIST          = "SELECT id, email, password, name, created, modified, last_enter FROM users"
	FIND          = "SELECT id, email, password, name, created, modified, last_enter FROM users WHERE id = $1"
	FIND_BY_EMAIL = "SELECT id, email, password, name, created, modified, last_enter FROM users WHERE email = $1"
	INSERT        = "INSERT INTO users (email, password, name, created, last_enter) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	UPDATE        = "UPDATE users SET email = $1, password = $2, name = $3, created = $4, last_enter = $6 WHERE id = $7"
	DELETE        = "DELETE FROM users WHERE id = $1"
)

func connection() *sql.DB {
	return db.Connection()
}

func scanRow(scaner db.RowScanner) *User {
	var user *User = new(User)
	err := scaner.Scan(&user.Id, &user.Email, &user.Password, &user.Name, &user.Created, &user.Modified, &user.LastEnter)
	if err != nil {
		logger.Errorf("Can't scan row: %s", err)
	}

	return user
}

func interateRows(rows *sql.Rows) []*User {
	users := make([]*User, 0)
	for rows.Next() {
		users = append(users, scanRow(rows))
	}

	return users
}

func List() []*User {
	rows, err := connection().Query(LIST)
	if err != nil {
		logger.Errorf("Database error while getting list of users: %s", err)
	}

	return interateRows(rows)
}

func Find(id int64) *User {
	return scanRow(connection().QueryRow(FIND, id))
}

func FindByEmail(email string) *User {
	return scanRow(connection().QueryRow(FIND_BY_EMAIL, email))
}

func (u *User) Subscription() *subscription.UserSubscription {
	return subscription.FindByUser(u.Id)
}

func (u *User) Save() {
	if u.Id != 0 {
		u.update()
	} else {
		u.insert()
	}
}

func (u *User) update() {
	_, err := connection().Exec(UPDATE, u.Email, u.Password, u.Name, u.Created, u.LastEnter, u.Id)
	if err != nil {
		logger.Errorf("Database error while updating user %d: %s", u.Id, err)
	}
}

func (u *User) insert() {
	err := connection().QueryRow(INSERT, u.Email, u.Password, u.Name, u.Created, u.LastEnter).Scan(&u.Id)
	if err != nil {
		logger.Errorf("Database error while inserting user: %s", err)
	}
}

func (u *User) Delete() {
	_, err := connection().Exec(DELETE, u.Id)
	if err != nil {
		logger.Errorf("Database error while deleting user %d: %s", u.Id, err)
	}
}

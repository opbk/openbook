package user

import "time"

type User struct {
	Id        int64
	Email     string
	Password  string
	Name      string
	Phone     string
	Created   time.Time
	Modified  time.Time
	LastEnter time.Time
}

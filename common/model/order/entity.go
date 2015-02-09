package order

import "time"

const (
	NEW        = "new"
	INPROGRESS = "inprogress"
	DELIVERED  = "delivered"
	ONHAND     = "onhand"
	RETURNED   = "returned"
)

type Order struct {
	Id        int64
	UserId    int64
	AddressId int64
	Status    string
	Comment   string
	Created   time.Time
	Modified  time.Time
}

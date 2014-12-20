package order

import "time"

const (
	NEW         = "new"
	IN_PROGRESS = "in_progress"
)

type Order struct {
	Id        int64     `json:"id"`
	UserId    int64     `json:"user_id"`
	AddressId int64     `json:"address_id"`
	Status    string    `json:"status"`
	Comment   string    `json:"comment"`
	Updated   time.Time `json:"updated"`
}

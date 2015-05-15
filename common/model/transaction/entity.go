package transaction

const (
	NEW       = "new"
	COMPLITED = "complited"
)

type Transaction struct {
	Id             int64
	UserId         int64
	SubscriptionId int64
	Payload        string
	Status         string
}

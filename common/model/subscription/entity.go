package subscription

type Subscription struct {
	Id    int64   `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

package price

type Price struct {
	Id     int64   `json:"id"`
	BookId int64   `json:"book_id"`
	Type   string  `json:"type"`
	Price  float64 `json:"price"`
}

package category

type Category struct {
	Id         int64  `json:"id"`
	CategoryId int64  `json:"category_id"`
	Name       string `json:"name"`
}

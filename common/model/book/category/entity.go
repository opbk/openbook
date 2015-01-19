package category

type Category struct {
	Id       int64  `json:"id"`
	ParentId int64  `json:"parent_id"`
	Path     string `json:"path"`
	Name     string `json:"name"`
	Books    int64  `json:"books"`
}

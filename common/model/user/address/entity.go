package address

type Address struct {
	Id      int64  `json:"id"`
	UserId  int64  `json:"user_id"`
	Address string `json:"address"`
	Comment string `json:"comment"`
}

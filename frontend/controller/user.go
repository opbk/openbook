package controller

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/opbk/openbook/common/model/book"
	"github.com/opbk/openbook/common/model/order"
	"github.com/opbk/openbook/common/web"
	"github.com/opbk/openbook/common/web/auth"
)

type UserController struct {
	FrontendController
	DefaultLimitPerPage int
}

func NewUserController() *UserController {
	c := new(UserController)
	c.DefaultLimitPerPage = 10
	return c
}

func (c *UserController) Routes(router *mux.Router) {
	router.HandleFunc("/user/me", auth.UserRequired(c.History))
	router.HandleFunc("/user/me/history", auth.UserRequired(c.History))
	router.HandleFunc("/user/me/wishlisth", auth.UserRequired(c.WishList))
}

func (c *UserController) History(rw http.ResponseWriter, req *http.Request) {
	request := web.NewRequest(req)

	limit := request.GetInt("l", c.DefaultLimitPerPage)
	offset := request.GetInt("f")
	orders := order.ListByUserWithLimit(c.getUser(req).Id, limit, offset)

	booksId := make([]int64, len(orders))
	for i, o := range orders {
		booksId[i] = o.BookId
	}

	c.ExecuteTemplate(rw, req, "history", map[string]interface{}{
		"orders": orders,
		"books":  book.MapById(booksId),
		"pagination": map[string]int{
			"total":  order.CountByUser(c.getUser(req).Id),
			"limit":  limit,
			"offset": offset,
		},
	})
}

func (c *UserController) WishList(rw http.ResponseWriter, req *http.Request) {
	c.ExecuteTemplate(rw, req, "wishlist", map[string]interface{}{})
}

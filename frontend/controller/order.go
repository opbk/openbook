package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/opbk/openbook/common/model/book"
	"github.com/opbk/openbook/common/model/order"
	"github.com/opbk/openbook/common/model/subscription"
	"github.com/opbk/openbook/common/web"
	"github.com/opbk/openbook/common/web/auth"
	"github.com/opbk/openbook/common/web/form/utils"
	"github.com/opbk/openbook/frontend/form"
)

type OrderController struct {
	FrontendController
}

func NewOrderController() *OrderController {
	return new(OrderController)
}

func (c *OrderController) Routes(router *mux.Router) {
	router.HandleFunc("/order", auth.UserRequired(c.Order)).Methods("POST")
	router.HandleFunc("/order/{bookId:[0-9]+}", auth.UserRequired(c.Order))
}

func (c *OrderController) Order(rw http.ResponseWriter, req *http.Request) {
	request := web.NewRequest(req)
	user := c.getUser(req)

	f := form.NewOrder()
	if req.Method == "POST" {
		if err := request.GetForm(f); err == nil && utils.IsValid(f) {
			user.Name = f.Name
			user.Phone = f.Phone
			user.Save()

			o := order.Order{
				UserId:    user.Id,
				AddressId: f.Address,
				Status:    order.NEW,
			}
			o.Save()
			order.AddOrderBook(o.Id, f.Book)

			if user.Subscription().Id == 0 && f.Subscription != 0 {
				http.Redirect(rw, req, fmt.Sprintf("/user/me/subscribe/%d", f.Subscription), http.StatusFound)
			} else {
				http.Redirect(rw, req, "/user/me/history", http.StatusFound)
			}
		}
	}

	book := book.Find(request.GetInt64("bookId", 0))
	if book == nil {
		http.NotFound(rw, req)
	}

	addresses := user.Addresses()
	subscriptions := subscription.List()

	f.Book = book.Id
	f.Name = user.Name
	f.Phone = user.Phone

	c.ExecuteTemplate(rw, req, "order", map[string]interface{}{
		"book":          book,
		"addresses":     addresses,
		"subscriptions": subscriptions,
		"form":          f,
	})
}

package controller

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/opbk/openbook/common/configuration"
	"github.com/opbk/openbook/common/mail"
	"github.com/opbk/openbook/common/model/author"
	"github.com/opbk/openbook/common/model/book"
	"github.com/opbk/openbook/common/model/order"
	"github.com/opbk/openbook/common/model/publisher"
	"github.com/opbk/openbook/common/model/subscription"
	"github.com/opbk/openbook/common/model/user"
	"github.com/opbk/openbook/common/model/user/address"
	"github.com/opbk/openbook/common/web"
	"github.com/opbk/openbook/common/web/auth"
	"github.com/opbk/openbook/common/web/form/utils"
	"github.com/opbk/openbook/frontend/form"
)

const (
	defaultLimitPerPage = 10
)

type OrderController struct {
	FrontendController
}

func NewOrderController() *OrderController {
	return new(OrderController)
}

func (c *OrderController) Routes(router *mux.Router) {
	router.HandleFunc("/order", auth.UserRequired(c.Order)).Methods("POST")
	router.HandleFunc("/order/{id:[0-9]+}/change", auth.UserRequired(c.Change)).Methods("GET")
	router.HandleFunc("/order/{id:[0-9]+}/return", auth.UserRequired(c.Return)).Methods("GET")
	router.HandleFunc("/order/{id:[0-9]+}", auth.UserRequired(c.Delete)).Methods("DELETE")
	router.HandleFunc("/order/book/{bookId:[0-9]+}", auth.UserRequired(c.Order))
	router.HandleFunc("/order/history", auth.UserRequired(c.History))
	router.HandleFunc("/order/history/{id:[0-9]+}", auth.UserRequired(c.History))
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
			c.orderNotifying(o.Id)

			if user.Subscription() == nil && f.Subscription != 0 {
				http.Redirect(rw, req, fmt.Sprintf("/user/me/subscribe/%d", f.Subscription), http.StatusFound)
			} else {
				http.Redirect(rw, req, "/order/history", http.StatusFound)
			}
			return
		}
	}
	fmt.Println(f.Errors)

	book := book.Find(request.GetInt64("bookId", 0))
	if book == nil {
		http.NotFound(rw, req)
	}

	authors := authorMapById(book.AuthorsId)
	publishers := publisherMapById([]int64{book.PublisherId})

	addresses := user.Addresses()
	subscriptions := subscription.List()

	f.Book = book.Id
	f.Name = user.Name
	f.Phone = user.Phone
	if len(addresses) > 0 {
		f.Address = addresses[0].Id
	}

	c.ExecuteTemplate(rw, req, "order", map[string]interface{}{
		"book":          book,
		"authors":       authors,
		"publishers":    publishers,
		"addresses":     addresses,
		"subscriptions": subscriptions,
		"form":          f,
	})
}

func (c *OrderController) History(rw http.ResponseWriter, req *http.Request) {
	request := web.NewRequest(req)

	status := request.GetString("s")
	limit := request.GetInt("l", defaultLimitPerPage)
	offset := request.GetInt("f")

	orders := make([]*order.Order, 0)
	if id := request.GetInt64("id"); id != 0 {
		orders = append(orders, order.Find(id))
	} else {
		orders = order.ListByUserAndStatusWithLimit(c.getUser(req).Id, status, limit, offset)
	}

	booksId := make([]int64, len(orders))
	for i, o := range orders {
		booksId[i] = o.BookId
	}

	books := make(map[int64]*book.Book)
	authors := make(map[int64]*author.Author)
	publishers := make(map[int64]*publisher.Publisher)

	if len(booksId) > 0 {
		books = book.MapById(booksId)
		authorsId := make([]int64, len(books))
		publishersId := make([]int64, len(books))
		for _, book := range books {
			authorsId = append(authorsId, book.AuthorsId...)
			publishersId = append(publishersId, book.PublisherId)
		}

		authors = authorMapById(authorsId)
		publishers = publisherMapById(publishersId)
	}

	addresses := make(map[int64]*address.Address)
	for _, address := range c.getUser(req).Addresses() {
		addresses[address.Id] = address
	}

	c.ExecuteTemplate(rw, req, "history", map[string]interface{}{
		"orders":    orders,
		"addresses": addresses,
		"books": map[string]interface{}{
			"books":      books,
			"authors":    authors,
			"publishers": publishers,
		},
		"pagination": map[string]int{
			"total":  order.CountByUserAndStatus(c.getUser(req).Id, status),
			"limit":  limit,
			"offset": offset,
		},
	})
}

func (c *OrderController) Change(rw http.ResponseWriter, req *http.Request) {
	request := web.NewRequest(req)
	if o := order.Find(request.GetInt64("id", 0)); o != nil {
		c.changeNotifying(o.Id)
		c.ExecuteTemplate(rw, req, "change", map[string]interface{}{})
		return
	}

	web.NotFound(rw)
}

func (c *OrderController) Return(rw http.ResponseWriter, req *http.Request) {
	request := web.NewRequest(req)
	if o := order.Find(request.GetInt64("id", 0)); o != nil {
		if o.Status == order.ONHAND && o.UserId == c.getUser(req).Id {
			o.Return()
			return
		}
	}

	web.NotFound(rw)
}

func (c *OrderController) Delete(rw http.ResponseWriter, req *http.Request) {
	request := web.NewRequest(req)
	if o := order.Find(request.GetInt64("id", 0)); o != nil {
		if o.Status == order.NEW && o.UserId == c.getUser(req).Id {
			o.Delete()
			return
		}
	}

	web.NotFound(rw)
}

func (c *OrderController) orderNotifying(orderId int64) {
	go func() {
		o := order.Find(orderId)
		book := book.Find(o.BookId)
		user := user.Find(o.UserId)
		address := address.Find(o.AddressId)
		date := time.Now().AddDate(0, 0, 1)

		templateFabric := func(name string) string {
			body := bytes.NewBuffer([]byte{})
			c.Template().ExecuteTemplate(body, name, map[string]interface{}{
				"user":    user,
				"book":    book,
				"address": address,
				"domain":  configuration.GetConfig().Main.Domain,
				"date":    date,
				"order":   o,
			})
			return body.String()
		}

		mail.SendTo(user.Email, "Информация о сделанном заказе", templateFabric("email_order_user"))
		mail.SendTo(configuration.GetConfig().Main.InfoEmail, "Поступление нового заказа", templateFabric("email_order_admin"))
	}()
}

func (c *OrderController) changeNotifying(orderId int64) {
	go func() {
		o := order.Find(orderId)
		book := book.Find(o.BookId)

		body := bytes.NewBuffer([]byte{})
		c.Template().ExecuteTemplate(body, "email_order_change", map[string]interface{}{
			"user":  user.Find(o.UserId),
			"order": o,
			"book":  book,
		})

		mail.SendTo(configuration.GetConfig().Main.InfoEmail, "Запрос на изменение заказа", body.String())
	}()
}

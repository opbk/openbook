package controller

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/opbk/openbook/common/web/auth"
)

type OrderController struct {
	FrontendController
}

func NewOrderController() *OrderController {
	return new(OrderController)
}

func (c *OrderController) Routes(router *mux.Router) {
	router.HandleFunc("/order", auth.UserRequired(c.Order))
}

func (c *OrderController) Order(rw http.ResponseWriter, req *http.Request) {
	c.ExecuteTemplate(rw, req, "order", map[string]interface{}{})
}

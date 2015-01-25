package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

type OrderController struct {
	FrontendController
}

func NewOrderController() *OrderController {
	return new(OrderController)
}

func (c *OrderController) Routes(router *mux.Router) {
	router.HandleFunc("/order", c.Order)
}

func (c *OrderController) Order(rw http.ResponseWriter, req *http.Request) {
	http.Redirect(rw, req, "/signup", http.StatusFound)
}

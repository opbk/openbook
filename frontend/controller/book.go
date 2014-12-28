package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	// "github.com/opbk/openbook/common/model/author"
	"github.com/opbk/openbook/common/model/book"
	"github.com/opbk/openbook/common/web"
)

type BookController struct {
	DefaultLimitPerPage int
}

func NewBookController() *BookController {
	c := new(BookController)
	c.DefaultLimitPerPage = 10
	return c
}

func (c *BookController) Routes(router *mux.Router) {
	router.HandleFunc("/search", c.ListByCategory)
}

func (c *BookController) Search(rw http.ResponseWriter, req *http.Request) {

}

package controller

import (
	"github.com/gorilla/mux"
	"net/http"

	"github.com/opbk/openbook/common/model/author"
	"github.com/opbk/openbook/common/web"
)

type AuthorController struct{}

func NewAuthorController() *AuthorController {
	return new(AuthorController)
}

func (c *AuthorController) Routes(router *mux.Router) {
	router.HandleFunc("/authors", c.List)
	router.HandleFunc("/authors/{id:[0-9]+}", c.Get)
}

func (c *AuthorController) List(rw http.ResponseWriter, req *http.Request) {
	web.WriteJson(rw, author.List())
}

func (c *AuthorController) Get(rw http.ResponseWriter, req *http.Request) {
	request := web.NewRequest(req)
	web.WriteJson(rw, author.Find(request.GetInt64("id")))
}

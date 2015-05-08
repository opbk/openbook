package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

type IndexController struct {
	FrontendController
}

func NewIndexController() *IndexController {
	return new(IndexController)
}

func (c *IndexController) Routes(router *mux.Router) {
	router.HandleFunc("/", c.Index)
	router.HandleFunc("/about", c.About)
}

func (c *IndexController) Index(rw http.ResponseWriter, req *http.Request) {
	http.Redirect(rw, req, "/search", http.StatusFound)
}

func (c *IndexController) About(rw http.ResponseWriter, req *http.Request) {
	c.ExecuteTemplate(rw, req, "about", map[string]interface{}{})
}

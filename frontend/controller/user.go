package controller

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/opbk/openbook/common/web/auth"
)

type UserController struct {
	FrontendController
}

func NewUserController() *UserController {
	return new(UserController)
}

func (c *UserController) Routes(router *mux.Router) {
	router.HandleFunc("/user/me", auth.UserRequired(c.History))
	router.HandleFunc("/user/me/history", auth.UserRequired(c.History))
	router.HandleFunc("/user/me/wishlisth", auth.UserRequired(c.WishList))
}

func (c *UserController) History(rw http.ResponseWriter, req *http.Request) {
	c.ExecuteTemplate(rw, req, "history", map[string]interface{}{})
}

func (c *UserController) WishList(rw http.ResponseWriter, req *http.Request) {
	c.ExecuteTemplate(rw, req, "wishlist", map[string]interface{}{})
}

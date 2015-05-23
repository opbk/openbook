package controller

import (
	"net/http"

	"github.com/gorilla/mux"

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
	router.HandleFunc("/user/me", auth.UserRequired(c.Me))
}

func (c *UserController) Me(rw http.ResponseWriter, req *http.Request) {

}

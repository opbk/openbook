package controller

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/opbk/openbook/common/model/user/address"
	"github.com/opbk/openbook/common/web"
	"github.com/opbk/openbook/common/web/auth"
)

type AddressController struct {
	FrontendController
}

func NewAddressController() *AddressController {
	return new(AddressController)
}

func (c *AddressController) Routes(router *mux.Router) {
	router.HandleFunc("/user/me/addresses", auth.UserRequired(c.Add)).Methods("POST")
}

func (c *AddressController) Add(rw http.ResponseWriter, req *http.Request) {
	request := web.NewRequest(req)

	a := new(address.Address)
	if err := request.GetObject(a); err != nil {
		web.BadRequest(rw, "Can't decode json object")
	}

	a.UserId = c.getUser(req).Id
	a.Save()
	web.WriteJson(rw, a)
}

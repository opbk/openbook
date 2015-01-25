package controller

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/opbk/openbook/common/model/user"
	"github.com/opbk/openbook/common/web"
)

type SignController struct {
	FrontendController
}

func NewSignController() *SignController {
	return new(SignController)
}

func (c *SignController) Routes(router *mux.Router) {
	router.HandleFunc("/signup", c.SignUp)
}

func (c *SignController) SignUp(rw http.ResponseWriter, req *http.Request) {
	request := web.NewRequest(req)

	var u *user.User
	if email := request.GetString("email"); email != "" {
		u = new(user.User)
		u.Email = email
		u.Password = user.GenPassword()
		u.Created = time.Now()
		u.Modified = time.Now()
		u.LastEnter = time.Now()
		u.Save()
	}

	c.Template().ExecuteTemplate(rw, "signup", map[string]interface{}{
		"user": u,
	})
}

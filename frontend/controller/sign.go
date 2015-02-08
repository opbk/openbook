package controller

import (
	"bytes"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/opbk/openbook/common/mail"
	"github.com/opbk/openbook/common/model/user"
	"github.com/opbk/openbook/common/web"
	"github.com/opbk/openbook/common/web/auth"
	"github.com/opbk/openbook/common/web/form/utils"
	"github.com/opbk/openbook/frontend/form"
)

type SignController struct {
	FrontendController
}

func NewSignController() *SignController {
	return new(SignController)
}

func (c *SignController) Routes(router *mux.Router) {
	router.HandleFunc("/signup", c.Sign)
	router.HandleFunc("/signout", auth.UserRequired(c.SignOut))
}

func (c *SignController) Sign(rw http.ResponseWriter, req *http.Request) {
	request := web.NewRequest(req)
	f := form.NewSignUp()
	f.From = request.GetString("from", "/")

	if req.Method == "POST" {
		if err := request.GetForm(f); err == nil && utils.IsValid(f) {
			var u *user.User
			if f.New {
				u = new(user.User)
				u.Email = f.Email
				u.Password = user.GenPassword()
				u.Created = time.Now()
				u.Modified = time.Now()
				u.LastEnter = time.Now()
				u.Save()
			} else {
				u = f.User
			}

			auth.Set(u, rw, req)
			http.Redirect(rw, req, f.From, http.StatusFound)
		}
	}

	c.ExecuteTemplate(rw, req, "signup", map[string]interface{}{
		"form": f,
	})
}

func (c *SignController) SignOut(rw http.ResponseWriter, req *http.Request) {
	request := web.NewRequest(req)
	auth.Delete(rw, req)

	http.Redirect(rw, req, request.GetString("from", "/"), http.StatusFound)
}

func (c *SignController) sendEmail(u *user.User) {
	go func() {
		body := bytes.NewBuffer([]byte{})
		c.Template().ExecuteTemplate(body, "email_signup", map[string]interface{}{
			"user": u,
		})

		mail.SendTo(u.Email, "Добро пожаловать в нашу библиотеку", body.String())
	}()
}

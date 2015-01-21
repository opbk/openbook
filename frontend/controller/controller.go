package controller

import (
	"html/template"
	"net/http"

	"github.com/gorilla/context"

	"github.com/opbk/openbook/common/model/user"
	"github.com/opbk/openbook/common/web"
)

type FrontendController struct {
	template *template.Template
}

func (c *FrontendController) getUser(req *http.Request) *user.User {
	return context.Get(req, "user").(*user.User)
}

func GetControllers() []web.Controller {
	return []web.Controller{
		NewBookController(),
	}
}

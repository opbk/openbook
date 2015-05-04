package controller

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/opbk/openbook/common/model/subscription"
	"github.com/opbk/openbook/common/web"
	"github.com/opbk/openbook/common/web/auth"
)

type SubscriptionController struct {
	FrontendController
}

func NewSubscriptionController() *SubscriptionController {
	return new(SubscriptionController)
}

func (c *SubscriptionController) Routes(router *mux.Router) {
	router.HandleFunc("/user/me/subscribe/{id:[0-9]+}", auth.UserRequired(c.subscribe))
}

func (c *SubscriptionController) subscribe(rw http.ResponseWriter, req *http.Request) {
	request := web.NewRequest(req)
	s := subscription.Find(request.GetInt64("id", 0))
	if s == nil {
		http.NotFound(rw, req)
	}

	c.ExecuteTemplate(rw, req, "subscribe", map[string]interface{}{
		"s": s,
	})
}

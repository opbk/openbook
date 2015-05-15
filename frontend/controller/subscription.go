package controller

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/opbk/openbook/common/configuration"
	"github.com/opbk/openbook/common/model/subscription"
	"github.com/opbk/openbook/common/model/transaction"
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
	subscription := subscription.Find(request.GetInt64("id", 0))
	if subscription == nil {
		http.NotFound(rw, req)
	}

	tx := transaction.NewTransaction(c.getUser(req).Id)
	tx.SubscriptionId = subscription.Id
	tx.Save()

	c.ExecuteTemplate(rw, req, "subscribe", map[string]interface{}{
		"s":           subscription,
		"transaction": tx,
		"yandex":      configuration.GetConfig().YandexMoney,
	})
}

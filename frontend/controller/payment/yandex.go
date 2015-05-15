package payment

import (
	"net/http"

	logger "github.com/cihub/seelog"
	"github.com/gorilla/mux"

	"github.com/opbk/openbook/common/configuration"
	"github.com/opbk/openbook/common/model/subscription"
	"github.com/opbk/openbook/common/model/transaction"
	"github.com/opbk/openbook/common/model/user"
	"github.com/opbk/openbook/common/web"
	"github.com/opbk/openbook/common/yandex"
	"github.com/opbk/openbook/frontend/controller"
)

type YandexMoneyController struct {
	controller.FrontendController
}

func NewYandexMoneyController() *YandexMoneyController {
	return new(YandexMoneyController)
}

func (c *YandexMoneyController) Routes(router *mux.Router) {
	router.HandleFunc("/payment/yandex/check", c.handleRequest(c.check))
	router.HandleFunc("/payment/yandex/payment", c.handleRequest(c.payment))
}

func (c *YandexMoneyController) check(ym *yandex.YandexMoney) interface{} {
	if !ym.IsValid() {
		logger.Errorf("Md5 of check request is invalid: %s", ym.Request())
		return ym.CheckResponse(yandex.AUTH_ERROR, "")
	}

	request := ym.Request()
	tx := transaction.Find(request.OrderNumber)
	if tx != nil {
		s := subscription.Find(tx.SubscriptionId)
		u := user.Find(tx.UserId)
		if u != nil && s != nil && s.Price == request.OrderSumAmount {
			return ym.CheckResponse(yandex.SUCCESS, "")
		}
	}

	return ym.CheckResponse(yandex.PAYMENT_ERROR, "")
}

func (c *YandexMoneyController) payment(ym *yandex.YandexMoney) interface{} {
	if !ym.IsValid() {
		logger.Errorf("Md5 of payment request is invalid: %s", ym.Request())
		return ym.PaymentResponse(yandex.AUTH_ERROR, "")
	}

	request := ym.Request()
	tx := transaction.Find(request.OrderNumber)
	s := subscription.Find(tx.SubscriptionId)
	u := user.Find(tx.UserId)
	u.Subscribe(s)
	tx.Complite()

	return ym.PaymentResponse(yandex.SUCCESS, "")
}

func (c *YandexMoneyController) handleRequest(f func(*yandex.YandexMoney) interface{}) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		config := configuration.GetConfig().YandexMoney
		ym := &yandex.YandexMoney{
			ShopId:       config.ShopId,
			Scid:         config.Scid,
			ShopPassword: config.ShopPassword,
		}
		ym.Parse(req)
		web.WriteXml(rw, f(ym))
	}
}

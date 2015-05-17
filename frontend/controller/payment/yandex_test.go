package payment

import (
	"time"

	"gopkg.in/check.v1"

	"github.com/opbk/openbook/common/db"
	"github.com/opbk/openbook/common/model/subscription"
	usubscription "github.com/opbk/openbook/common/model/user/subscription"
	"github.com/opbk/openbook/common/web/webtest"
)

type YandexControllerTestSuite struct {
	webtest.UserControllerTestSuit
}

func (s *YandexControllerTestSuite) SetUpSuite(c *check.C) {
	router := s.GetRouter()
	controller := NewYandexMoneyController()
	controller.Routes(router)
	s.ControllerTestSuit.SetUpSuite(c)
}

func (s *YandexControllerTestSuite) SetUpTest(c *check.C) {
	db.Connection().Exec("ALTER SEQUENCE auto_id_transactions RESTART WITH 1")
	db.Connection().Exec("TRUNCATE transactions")
	db.Connection().Exec("ALTER SEQUENCE auto_id_subscriptions RESTART WITH 1")
	db.Connection().Exec("TRUNCATE subscriptions")
	db.Connection().Exec("TRUNCATE user_subscriptions")

	subscr := &subscription.Subscription{0, "Basic subscription", "Basic one month subscription", 490, true}
	subscr.Save()
	(&usubscription.UserSubscription{*subscr, 1, time.Now()}).Insert()
}

func (s *YandexControllerTestSuite) TearDownSuite(c *check.C) {
	db.Connection().Exec("ALTER SEQUENCE auto_id_transactions RESTART WITH 1")
	db.Connection().Exec("TRUNCATE transactions")
	db.Connection().Exec("ALTER SEQUENCE auto_id_subscriptions RESTART WITH 1")
	db.Connection().Exec("TRUNCATE subscriptions")
	db.Connection().Exec("TRUNCATE user_subscriptions")

	s.ControllerTestSuit.TearDownSuite(c)
}

// customerNumber=2362&sumCurrency=10643&&shopSumCurrencyPaycash=10643&orderSumAmount=40.00&shopId=35545&action=checkOrder&orderCreatedDatetime=2015-04-02T14%3A12%3A38.798%2B03%3A00&shopSumBankPaycash=1003&requestDatetime=2015-04-02T14%3A12%3A46.151%2B03%3A00&shopSumAmount=38.00&orderSumCurrencyPaycash=10643&orderSumBankPaycash=1003&invoiceId=2000000440348&paymentType=PC&&paymentPayerCode=4100322344779&orderNumber=75f105ea-998c-4600-8f0d-0d87b9d4a90a&md5=10DF35BF3E9225A6169B39090CC4C01E

func (s *YandexControllerTestSuite) TestCheckInvalidSignature(c *check.C) {
	res, _ := s.Post("/payment/yandex/check", "")
	c.Assert(res, check.Matches, "(?Us).*checkOrderResponse.*code=\"1\".*")
}

func (s *YandexControllerTestSuite) TestPaymentInvalidSignature(c *check.C) {
	res, _ := s.Post("/payment/yandex/payment", "")
	c.Assert(res, check.Matches, "(?Us).*paymentAvisoResponse.*code=\"1\".*")
}

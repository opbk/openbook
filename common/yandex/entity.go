package yandex

const (
	SUCCESS       = 0
	AUTH_ERROR    = 1
	PAYMENT_ERROR = 100
	PARSE_ERROR   = 200
)

type YandexMoney struct {
	ShopId       string
	Scid         string
	ShopPassword string
	request      *Request
}

type Request struct {
	RequestDatetime         string  `schema:"requestDatetime"`
	Action                  string  `schema:"action"`
	Md5                     string  `schema:"md5"`
	ShopId                  int64   `schema:"shopId"`
	ShopArticleId           int64   `schema:"shopArticleId"`
	InvoiceId               int64   `schema:"invoiceId"`
	OrderNumber             int64   `schema:"orderNumber"`
	CustomerNumber          string  `schema:"customerNumber"`
	OrderCreatedDatetime    string  `schema:"orderCreatedDatetime"`
	OrderSumAmount          float64 `schema:"orderSumAmount"`
	OrderSumCurrencyPaycash string  `schema:"orderSumCurrencyPaycash"`
	OrderSumBankPaycash     string  `schema:"orderSumBankPaycash"`
	ShopSumAmount           string  `schema:"shopSumAmount"`
	ShopSumCurrencyPaycash  string  `schema:"shopSumCurrencyPaycash"`
	ShopSumBankPaycash      string  `schema:"shopSumBankPaycash"`
	PaymentPayerCode        string  `schema:"paymentPayerCode"`
	PaymentType             string  `schema:"paymentType"`
	PaymentDatetime         string  `schema:"paymentDatetime"`
}

type Response struct {
	PerformedDatetime string  `xml:"performedDatetime,attr"`
	Code              int     `xml:"code,attr"`
	ShopId            int64   `xml:"shopId,attr"`
	InvoiceId         int64   `xml:"invoiceId,attr"`
	OrderSumAmount    float64 `xml:"orderSumAmount,attr"`
	Message           string  `xml:"message,attr,omitempty"`
	TechMessage       string  `xml:"techMessage,attr,omitempty"`
}

type checkOrderResponse struct {
	Response
}

type paymentAvisoResponse struct {
	Response
}

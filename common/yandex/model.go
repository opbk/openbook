package yandex

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/schema"
)

func (ym *YandexMoney) Parse(req *http.Request) error {
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	req.ParseForm()

	ym.request = new(Request)
	return decoder.Decode(ym.request, req.PostForm)
}

func (ym *YandexMoney) IsValid() bool {
	_sig := fmt.Sprintf("%s;%.2f;%s;%s;%d;%d;%s;%s",
		ym.request.Action, ym.request.OrderSumAmount, ym.request.OrderSumCurrencyPaycash, ym.request.OrderSumBankPaycash, ym.request.ShopId,
		ym.request.InvoiceId, ym.request.CustomerNumber, ym.ShopPassword)

	hasher := md5.New()
	hasher.Write([]byte(_sig))
	return ym.request.Md5 == strings.ToUpper(hex.EncodeToString(hasher.Sum(nil)))
}

func (ym *YandexMoney) Request() *Request {
	return ym.request
}

func (ym *YandexMoney) CheckResponse(code int, message string) checkOrderResponse {
	return checkOrderResponse{Response{
		PerformedDatetime: time.Now().Format("2006-01-02T15:04:05.000-07:00"),
		Code:              code,
		ShopId:            ym.request.ShopId,
		InvoiceId:         ym.request.InvoiceId,
		OrderSumAmount:    ym.request.OrderSumAmount,
		Message:           message,
	}}
}

func (ym *YandexMoney) PaymentResponse(code int, message string) paymentAvisoResponse {
	return paymentAvisoResponse{Response{
		PerformedDatetime: time.Now().Format("2006-01-02T15:04:05.000-07:00"),
		Code:              code,
		ShopId:            ym.request.ShopId,
		InvoiceId:         ym.request.InvoiceId,
		OrderSumAmount:    ym.request.OrderSumAmount,
		Message:           message,
	}}
}

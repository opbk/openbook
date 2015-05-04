package form

import (
	"errors"

	"github.com/opbk/openbook/common/model/book"
	"github.com/opbk/openbook/common/model/subscription"
	"github.com/opbk/openbook/common/model/user/address"
	"github.com/opbk/openbook/common/web/form"
)

type Order struct {
	form.Form
	Book         int64
	Name         string
	Phone        string
	Address      int64
	Subscription int64
}

func NewOrder() *Order {
	f := new(Order)
	f.Errors = make(map[string]error)
	return f
}

func (s *Order) Validate() {
	s.CheckIfEmpty("Name", s.Name, "Укажите как к Вам обращаться")
	s.CheckIfEmpty("Phone", s.Phone, "Укажите телефон, что бы мы могли связаться с Вами")

	book := book.Find(s.Book)
	if book == nil {
		s.Errors["Book"] = errors.New("Выбранная книга не найдена")
	}

	address := address.Find(s.Address)
	if address == nil {
		s.Errors["Address"] = errors.New("Выбранный адрес не найден")
	}

	subscription := subscription.Find(s.Subscription)
	if subscription == nil {
		s.Errors["Subscription"] = errors.New("Выбранная модель подписки не найдена")
	}
}

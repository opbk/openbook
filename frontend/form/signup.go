package form

import (
	"errors"
	"fmt"

	"github.com/opbk/openbook/common/model/user"
	"github.com/opbk/openbook/common/web/form"
)

type SighUp struct {
	form.Form
	Email    string
	Password string
	New      bool
	User     *user.User
	From     string
}

func NewSignUp() *SighUp {
	f := new(SighUp)
	f.Errors = make(map[string]error)
	return f
}

func (s *SighUp) Validate() {
	if !s.CheckIfEmpty("Email", s.Email, "Укажите ваш email адрес") {
		s.User = user.FindByEmail(s.Email)
		if s.User != nil && s.New {
			s.Errors["Email"] = errors.New(fmt.Sprintf("Пользователь %s уже существует", s.Email))
		}
	}

	if !s.New {
		if !s.CheckIfEmpty("Password", s.Password, "Укажите ваш пароль") {
			s.User = user.FindByEmail(s.Email)
			if s.User != nil && s.User.Password != s.Password {
				s.Errors["Password"] = errors.New(fmt.Sprintf("Указанный пароль для пользователя %s не подходит", s.Email))
			} else if s.User == nil {
				s.Errors["Email"] = errors.New(fmt.Sprintf("Пользователь %s не наден", s.Email))
			}
		}
	}
}

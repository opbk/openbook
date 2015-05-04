package form

import (
	"time"

	"gopkg.in/check.v1"

	"github.com/opbk/openbook/common/configuration"
	"github.com/opbk/openbook/common/db"
	"github.com/opbk/openbook/common/model/user"
	"github.com/opbk/openbook/common/web/form/utils"
)

type SignupTestSuit struct{}

var _ = check.Suite(new(SignupTestSuit))

func (s *SignupTestSuit) SetUpSuite(c *check.C) {
	config := configuration.LoadConfiguration("")
	db.InitDbConnection(config.Db.Driver, config.Db.Connection)
}

func (s *SignupTestSuit) SetUpTest(c *check.C) {
	db.Connection().Exec("ALTER SEQUENCE auto_id_users RESTART WITH 1")
	db.Connection().Exec("TRUNCATE users")

	(&user.User{0, "netw00rk@gmail.com", "123456", "A. Koklin", "89161234567", time.Now(), time.Now(), time.Now()}).Save()
}

func (s *SignupTestSuit) TearDownSuite(c *check.C) {
	db.Connection().Exec("ALTER SEQUENCE auto_id_users RESTART WITH 1")
	db.Connection().Exec("TRUNCATE users")
}

func (s *SignupTestSuit) TestNewWithEmptyEmail(c *check.C) {
	f := NewSignUp()
	f.New = true

	c.Assert(utils.IsValid(f), check.Equals, false)
	c.Assert(f.Error("Email"), check.Equals, "Укажите ваш email адрес")
	c.Assert(f.Error("Password"), check.Equals, "")
}

func (s *SignupTestSuit) TestWithEmptyPassword(c *check.C) {
	f := NewSignUp()
	f.New = false

	c.Assert(utils.IsValid(f), check.Equals, false)
	c.Assert(f.Error("Password"), check.Equals, "Укажите ваш пароль")
}

func (s *SignupTestSuit) TestNewWithExistUser(c *check.C) {
	f := NewSignUp()
	f.New = true
	f.Email = "netw00rk@gmail.com"

	c.Assert(utils.IsValid(f), check.Equals, false)
	c.Assert(f.Error("Email"), check.Equals, "Пользователь netw00rk@gmail.com уже существует")
}

func (s *SignupTestSuit) TestExistWithIncorectPassword(c *check.C) {
	f := NewSignUp()
	f.Email = "netw00rk@gmail.com"
	f.Password = "1234567"

	c.Assert(utils.IsValid(f), check.Equals, false)
	c.Assert(f.Error("Password"), check.Equals, "Указанный пароль для пользователя netw00rk@gmail.com не подходит")
}

func (s *SignupTestSuit) TestValidWithUser(c *check.C) {
	f := NewSignUp()
	f.Email = "netw00rk@gmail.com"
	f.Password = "123456"

	c.Assert(utils.IsValid(f), check.Equals, true)
	c.Assert(f.User.Id, check.Equals, int64(1))
}

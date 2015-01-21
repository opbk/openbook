package user

import (
	"testing"
	"time"

	"gopkg.in/check.v1"

	"github.com/opbk/openbook/common/configuration"
	"github.com/opbk/openbook/common/db"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type TestSuit struct{}

var _ = check.Suite(new(TestSuit))

func (s *TestSuit) SetUpSuite(c *check.C) {
	config := configuration.LoadConfiguration("")
	db.InitDbConnection(config.Db.Driver, config.Db.Connection)
}

func (s *TestSuit) SetUpTest(c *check.C) {
	connection().Exec("ALTER SEQUENCE auto_id_users RESTART WITH 1")
	connection().Exec("TRUNCATE users")

	(&User{0, "netw00rk@gmail.com", "123456", "A. Koklin", time.Now(), time.Now(), time.Now()}).Save()
}

func (s *TestSuit) TearDownSuite(c *check.C) {
	connection().Exec("ALTER SEQUENCE auto_id_users RESTART WITH 1")
	connection().Exec("TRUNCATE users")
}

func (s *TestSuit) TestFind(c *check.C) {
	u := Find(1)
	c.Assert(u.Email, check.Equals, "netw00rk@gmail.com")
	c.Assert(u.Password, check.Equals, "123456")
	c.Assert(u.Name, check.Equals, "A. Koklin")
}

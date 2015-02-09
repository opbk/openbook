package address

import (
	"testing"

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
	connection().Exec("ALTER SEQUENCE auto_id_addresses RESTART WITH 1")
	connection().Exec("TRUNCATE addresses")

	(&Address{0, 1, "Ул. Лизюкова, д. 21, кв. 41", ""}).Save()
	(&Address{0, 1, "Проспект Печкина, д. 30, офис 501", "Перед приходом позвонить"}).Save()
}

func (s *TestSuit) TearDownSuite(c *check.C) {
	connection().Exec("ALTER SEQUENCE auto_id_addresses RESTART WITH 1")
	connection().Exec("TRUNCATE addresses")
}

func (s *TestSuit) TestFind(c *check.C) {
	a := Find(2)
	c.Assert(a.UserId, check.Equals, int64(1))
	c.Assert(a.Address, check.Equals, "Проспект Печкина, д. 30, офис 501")
	c.Assert(a.Comment, check.Equals, "Перед приходом позвонить")
}

func (s *TestSuit) TestListByUser(c *check.C) {
	as := ListByUser(1)
	c.Assert(len(as), check.Equals, 2)
}

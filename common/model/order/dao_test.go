package order

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
	connection().Exec("ALTER SEQUENCE auto_id_orders RESTART WITH 1")
	connection().Exec("TRUNCATE orders")

	(&Order{0, 1, 1, NEW, "", time.Now(), time.Now()}).Save()
	(&Order{0, 1, 2, INPROGRESS, "", time.Now(), time.Now()}).Save()
}

func (s *TestSuit) TearDownSuite(c *check.C) {
	connection().Exec("ALTER SEQUENCE auto_id_orders RESTART WITH 1")
	connection().Exec("TRUNCATE orders")
}

func (s *TestSuit) TestFind(c *check.C) {
	o := Find(1)
	c.Assert(o.UserId, check.Equals, int64(1))
	c.Assert(o.AddressId, check.Equals, int64(1))
	c.Assert(o.Status, check.Equals, NEW)
}

func (s *TestSuit) TestListByUser(c *check.C) {
	os := ListByUser(1)
	c.Assert(len(os), check.Equals, 2)
}

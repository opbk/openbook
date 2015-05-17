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
	connection().Exec("TRUNCATE book_orders")

	(&Order{0, 0, 1, 1, NEW, "", time.Now(), time.Now()}).Save()
	(&Order{0, 0, 1, 2, INPROGRESS, "", time.Now(), time.Now()}).Save()
	AddOrderBook(1, 1)
}

func (s *TestSuit) TearDownSuite(c *check.C) {
	connection().Exec("ALTER SEQUENCE auto_id_orders RESTART WITH 1")
	connection().Exec("TRUNCATE orders")
	connection().Exec("TRUNCATE book_orders")
}

func (s *TestSuit) TestFind(c *check.C) {
	o := Find(1)
	c.Assert(o.UserId, check.Equals, int64(1))
	c.Assert(o.BookId, check.Equals, int64(1))
	c.Assert(o.AddressId, check.Equals, int64(1))
	c.Assert(o.Status, check.Equals, NEW)
}

func (s *TestSuit) TestListByUserWithLimit(c *check.C) {
	orders := ListByUserAndStatusWithLimit(1, "", 10, 0)
	c.Assert(orders[0].Id, check.Equals, int64(2))
	c.Assert(orders[1].Id, check.Equals, int64(1))
	c.Assert(len(orders), check.Equals, 2)
}

func (s *TestSuit) TestListByUserAndStatusWithLimit(c *check.C) {
	orders := ListByUserAndStatusWithLimit(1, INPROGRESS, 10, 0)
	c.Assert(orders[0].Id, check.Equals, int64(2))
	c.Assert(len(orders), check.Equals, 1)
}

func (s *TestSuit) TestCountByUser(c *check.C) {
	count := CountByUserAndStatus(1, "")
	c.Assert(count, check.Equals, 2)
}

func (s *TestSuit) TestCountByUserAndStatus(c *check.C) {
	count := CountByUserAndStatus(1, INPROGRESS)
	c.Assert(count, check.Equals, 1)
}

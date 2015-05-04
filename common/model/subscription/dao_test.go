package subscription

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
	connection().Exec("ALTER SEQUENCE auto_id_subscriptions RESTART WITH 1")
	connection().Exec("TRUNCATE subscriptions")

	(&Subscription{0, "Basic subscription", "Basic one month subscription", 490, true}).Save()
	(&Subscription{0, "Vip subscription", "Subscription with one free delivery every month", 690, false}).Save()
}

func (s *TestSuit) TearDownSuite(c *check.C) {
	connection().Exec("ALTER SEQUENCE auto_id_subscriptions RESTART WITH 1")
	connection().Exec("TRUNCATE subscriptions")
}

func (s *TestSuit) TestFind(c *check.C) {
	sub := Find(1)
	c.Assert(sub.Name, check.Equals, "Basic subscription")
	c.Assert(sub.Description, check.Equals, "Basic one month subscription")
	c.Assert(sub.Price, check.Equals, float64(490))
	c.Assert(sub.Enabled, check.Equals, true)
}

func (s *TestSuit) TestList(c *check.C) {
	subs := List()
	c.Assert(len(subs), check.Equals, 2)
}

func (s *TestSuit) TestListEnabled(c *check.C) {
	subs := ListEnabled()
	c.Assert(len(subs), check.Equals, 1)
}

package subscription

import (
	"testing"
	"time"

	"gopkg.in/check.v1"

	"github.com/opbk/openbook/common/configuration"
	"github.com/opbk/openbook/common/db"
	"github.com/opbk/openbook/common/model/subscription"
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
	connection().Exec("TRUNCATE user_subscriptions")

	sub := &subscription.Subscription{0, "Basic subscription", "Basic one month subscription", 490, true}
	sub.Save()

	(&UserSubscription{*sub, 1, time.Now()}).Insert()
}

func (s *TestSuit) TearDownSuite(c *check.C) {
	connection().Exec("ALTER SEQUENCE auto_id_subscriptions RESTART WITH 1")
	connection().Exec("TRUNCATE subscriptions")
	connection().Exec("TRUNCATE user_subscriptions")
}

func (s *TestSuit) TestFindByUser(c *check.C) {
	sub := FindByUser(1)

	c.Assert(sub.Name, check.Equals, "Basic subscription")
	c.Assert(sub.Description, check.Equals, "Basic one month subscription")
	c.Assert(sub.Price, check.Equals, float64(490))
	c.Assert(sub.Expiration, check.NotNil)
}

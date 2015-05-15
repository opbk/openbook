package transaction

import (
	"testing"

	"gopkg.in/check.v1"

	"github.com/opbk/openbook/common/configuration"
	"github.com/opbk/openbook/common/db"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type TestSuite struct{}

var _ = check.Suite(new(TestSuite))

func (s *TestSuite) SetUpSuite(c *check.C) {
	config := configuration.LoadConfiguration("")
	db.InitDbConnection(config.Db.Driver, config.Db.Connection)
}

func (s *TestSuite) SetUpTest(c *check.C) {
	connection().Exec("ALTER SEQUENCE auto_id_transactions RESTART WITH 1")
	connection().Exec("TRUNCATE transactions")

	tx := NewTransaction(100)
	tx.SubscriptionId = 50
	tx.Payload = "request_payload"
	tx.Save()
}

func (s *TestSuite) TearDownSuite(c *check.C) {
	connection().Exec("ALTER SEQUENCE auto_id_transactions RESTART WITH 1")
	connection().Exec("TRUNCATE transactions")
}

func (s *TestSuite) TestFind(c *check.C) {
	tx := Find(1)
	c.Assert(tx.UserId, check.Equals, int64(100))
	c.Assert(tx.SubscriptionId, check.Equals, int64(50))
	c.Assert(tx.Payload, check.Equals, "request_payload")
	c.Assert(tx.Status, check.Equals, NEW)
}

func (s *TestSuite) TestGetNewTransactionOrCreate(c *check.C) {
	tx := NewTransaction(100)
	c.Assert(tx.Id, check.Equals, int64(1))
	tx.Complite()

	tx = NewTransaction(100)
	c.Assert(tx.Id, check.Equals, int64(0))
}

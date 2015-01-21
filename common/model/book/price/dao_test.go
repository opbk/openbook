package price

import (
	"testing"

	"gopkg.in/check.v1"

	"github.com/opbk/openbook/common/configuration"
	"github.com/opbk/openbook/common/db"
	"github.com/opbk/openbook/common/model/book/price/pricetype"
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
	connection().Exec("ALTER SEQUENCE auto_id_prices RESTART WITH 1")
	connection().Exec("TRUNCATE prices")
	connection().Exec("TRUNCATE book_prices")

	(&pricetype.Type{0, "retail", "Розничная цена"}).Save()
	(&pricetype.Type{0, "month", "Аренда на месяц"}).Save()
	AddBookPrice(1, 1, 1000)
	AddBookPrice(1, 2, 300)
}

func (s *TestSuit) TearDownSuite(c *check.C) {
	connection().Exec("ALTER SEQUENCE auto_id_prices RESTART WITH 1")
	connection().Exec("TRUNCATE prices")
	connection().Exec("TRUNCATE book_prices")
}

func (s *TestSuit) TestMapByBookId(c *check.C) {
	prices := MapByBookId(1)

	c.Assert(prices["retail"].Price, check.Equals, 1000.0)
	c.Assert(prices["month"].Price, check.Equals, 300.0)
}

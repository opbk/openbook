package utils

import (
	"testing"
	"time"

	"gopkg.in/check.v1"

	"github.com/opbk/openbook/common/configuration"
	"github.com/opbk/openbook/common/db"

	"github.com/opbk/openbook/common/model/book"
	"github.com/opbk/openbook/common/model/book/category"
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
	connection().Exec("ALTER SEQUENCE auto_id_books RESTART WITH 1")
	connection().Exec("ALTER SEQUENCE auto_id_categories RESTART WITH 1")
	connection().Exec("TRUNCATE books")
	connection().Exec("TRUNCATE categories")
	connection().Exec("TRUNCATE book_categories")

	(&category.Category{0, 0, "", "Fiction", 0}).Save()
	(&category.Category{0, 1, ">1", "SF", 0}).Save()
	(&category.Category{0, 1, ">1", "Other", 0}).Save()
	(&category.Category{0, 2, ">1>2", "Cosmos", 0}).Save()

	(&book.Book{0, "The Martian", 600, "en", "", "Six days ago, astronaut Mark Watney became one of the first people to walk on Mars.", "", "", time.Date(2013, time.March, 22, 0, 0, 0, 0, time.Local), time.Now(), 0, 1, []int64{1}, []int64{500}}).Save()

	book.AddBookToCategory(1, 4)

}

func (s *TestSuit) TearDownSuite(c *check.C) {
	connection().Exec("ALTER SEQUENCE auto_id_books RESTART WITH 1")
	connection().Exec("ALTER SEQUENCE auto_id_categories RESTART WITH 1")
	connection().Exec("TRUNCATE books")
	connection().Exec("TRUNCATE categories")
	connection().Exec("TRUNCATE book_categories")

}

func (s *TestSuit) TestFixCategories(c *check.C) {
	var test_book = book.Find(int64(1))

	FixCategories(test_book)

	var categories = GetBookCategories(test_book)

	c.Check(categories[0].Id, check.Equals, int64(4))
	c.Check(categories[1].Id, check.Equals, int64(1))
	c.Check(categories[2].Id, check.Equals, int64(2))

}

func (s *TestSuit) TestDeleteConnections(c *check.C) {
	var test_book = book.Find(int64(1))

	DeleteConnections(test_book, category.Find(4))

}

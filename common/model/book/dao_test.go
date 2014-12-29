package book

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
	connection().Exec("ALTER SEQUENCE auto_id_books RESTART WITH 1")
	connection().Exec("TRUNCATE books")
	connection().Exec("TRUNCATE book_categories")
	connection().Exec("TRUNCATE author_books")

	(&Book{0, "The Martian", "Six days ago, astronaut Mark Watney became one of the first people to walk on Mars.", time.Date(2013, time.March, 22, 0, 0, 0, 0, time.Local), time.Now()}).Save()
	(&Book{0, "The Egg", "You were on your way home when you died.It was a car accident.", time.Date(2014, time.November, 26, 0, 0, 0, 0, time.Local), time.Now()}).Save()
	(&Book{0, "Inside the Doomsday Machine", "Who understood the risk inherent in the assumption...", time.Date(2010, time.March, 15, 0, 0, 0, 0, time.Local), time.Now()}).Save()
	(&Book{0, "The Hobbit", "Like every other hobbit, Bilbo Baggins likes nothing...", time.Date(2012, time.September, 10, 0, 0, 0, 0, time.Local), time.Now()}).Save()

	AddBookToCategory(1, 1)
	AddBookToCategory(2, 1)
	AddBookToCategory(3, 2)
	AddBookToCategory(4, 1)

	AddBookToAuthor(1, 1)
	AddBookToAuthor(2, 1)
	AddBookToAuthor(3, 2)
	AddBookToAuthor(4, 3)
}

func (s *TestSuit) TestIterateRows(c *check.C) {
	book := List(1, 0)[0]
	c.Assert(book.Id, check.Equals, int64(1))
	c.Assert(book.Title, check.Equals, "The Martian")
	c.Assert(book.Description, check.Equals, "Six days ago, astronaut Mark Watney became one of the first people to walk on Mars.")
	c.Assert(book.Release.Format("2006-01-02"), check.Equals, "2013-03-22")
	c.Assert(book.Created.Format("2006-01-02"), check.Equals, time.Now().Format("2006-01-02"))
}

func (s *TestSuit) TearDownSuite(c *check.C) {
	connection().Exec("ALTER SEQUENCE auto_id_books RESTART WITH 1")
	connection().Exec("TRUNCATE books")
	connection().Exec("TRUNCATE book_categories")
	connection().Exec("TRUNCATE author_books")
}

func (s *TestSuit) TestSearchEmptyWithLimit(c *check.C) {
	books := Search(map[string]interface{}{}, 3, 0)
	c.Assert(len(books), check.Equals, 3)
	c.Assert(books[0].Id, check.Equals, int64(4))
	c.Assert(books[1].Id, check.Equals, int64(3))
	c.Assert(books[2].Id, check.Equals, int64(2))
}

func (s *TestSuit) TestSearchEmptyWithLimitAndOffset(c *check.C) {
	books := Search(map[string]interface{}{}, 3, 3)
	c.Assert(len(books), check.Equals, 1)
	c.Assert(books[0].Id, check.Equals, int64(1))
}

func (s *TestSuit) TestSearchOnlyWithCategory(c *check.C) {
	books := Search(map[string]interface{}{"categories": []int64{1}}, 20, 0)
	c.Assert(len(books), check.Equals, 3)
	c.Assert(books[0].Id, check.Equals, int64(4))
	c.Assert(books[1].Id, check.Equals, int64(2))
	c.Assert(books[2].Id, check.Equals, int64(1))
}

func (s *TestSuit) TestSearchOnlyWithAuthor(c *check.C) {
	books := Search(map[string]interface{}{"author": 1}, 20, 0)
	c.Assert(len(books), check.Equals, 2)
	c.Assert(books[0].Id, check.Equals, int64(2))
	c.Assert(books[1].Id, check.Equals, int64(1))
}

func (s *TestSuit) TestSearchWithAuthorAndCategory(c *check.C) {
	books := Search(map[string]interface{}{"categories": []int64{1}, "author": 3}, 20, 0)
	c.Assert(len(books), check.Equals, 1)
	c.Assert(books[0].Id, check.Equals, int64(4))
}

func (s *TestSuit) TestSearchOnlyWithRealeas(c *check.C) {
	books := Search(map[string]interface{}{"release": "2013-01-01"}, 20, 0)
	c.Assert(len(books), check.Equals, 2)
	c.Assert(books[0].Id, check.Equals, int64(2))
	c.Assert(books[1].Id, check.Equals, int64(1))
}

func (s *TestSuit) TestSearchWithAuthorAndCategoryAndRelease(c *check.C) {
	books := Search(map[string]interface{}{"categories": []int64{1}, "author": 1, "release": "2014-01-01"}, 20, 0)
	c.Assert(len(books), check.Equals, 1)
	c.Assert(books[0].Id, check.Equals, int64(2))
}

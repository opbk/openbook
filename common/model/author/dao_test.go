package author

import (
	"testing"
	"time"

	"gopkg.in/check.v1"

	"github.com/opbk/openbook/common/configuration"
	"github.com/opbk/openbook/common/db"
	"github.com/opbk/openbook/common/model/book"
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
	connection().Exec("ALTER SEQUENCE auto_id_authors RESTART WITH 1")
	connection().Exec("ALTER SEQUENCE auto_id_books RESTART WITH 1")
	connection().Exec("TRUNCATE authors")
	connection().Exec("TRUNCATE books")
	connection().Exec("TRUNCATE book_categories")
	connection().Exec("TRUNCATE author_books")

	(&Author{0, "Andy Weir", "", 2}).Save()
	(&Author{0, "Michael Lewis", "", 1}).Save()
	(&Author{0, "John R. R. Tolkien", "", 1}).Save()

	(&book.Book{0, "The Martian", 600, "en", "", "Six days ago, astronaut Mark Watney became one of the first people to walk on Mars.", "", "", time.Date(2013, time.March, 22, 0, 0, 0, 0, time.Local), time.Now(), 0, 1, []int64{1}, []int64{500}}).Save()
	(&book.Book{0, "The Egg", 500, "en", "", "You were on your way home when you died.It was a car accident.", "", "", time.Date(2014, time.November, 26, 0, 0, 0, 0, time.Local), time.Now(), 0, 1, []int64{1}, []int64{1}}).Save()
	(&book.Book{0, "Inside the Doomsday Machine", 400, "en", "", "Who understood the risk inherent in the assumption...", "", "", time.Date(2010, time.March, 15, 0, 0, 0, 0, time.Local), time.Now(), 0, 2, []int64{2}, []int64{2}}).Save()
	(&book.Book{0, "The Hobbit", 700, "en", "", "Like every other hobbit, Bilbo Baggins likes nothing...", "", "", time.Date(2012, time.September, 10, 0, 0, 0, 0, time.Local), time.Now(), 0, 1, []int64{1}, []int64{3}}).Save()

	book.AddBookToCategory(1, 1)
	book.AddBookToCategory(1, 3)
	book.AddBookToCategory(2, 1)
	book.AddBookToCategory(3, 2)
	book.AddBookToCategory(4, 1)

	book.AddBookToAuthor(1, 1)
	book.AddBookToAuthor(2, 1)
	book.AddBookToAuthor(3, 2)
	book.AddBookToAuthor(4, 3)
}

func (s *TestSuit) TearDownSuite(c *check.C) {
	connection().Exec("ALTER SEQUENCE auto_id_authors RESTART WITH 1")
	connection().Exec("ALTER SEQUENCE auto_id_books RESTART WITH 1")
	connection().Exec("TRUNCATE authors")
	connection().Exec("TRUNCATE books")
	connection().Exec("TRUNCATE book_categories")
	connection().Exec("TRUNCATE author_books")
}

func (s *TestSuit) TestMapById(c *check.C) {
	authors := MapById([]int64{1, 3})
	c.Assert(authors[1].Name, check.Equals, "Andy Weir")
	c.Assert(authors[3].Name, check.Equals, "John R. R. Tolkien")
}

func (s *TestSuit) TestSearchEmpty(c *check.C) {
	authors := Search(map[string]interface{}{})

	c.Assert(authors[0].Id, check.Equals, int64(1))
	c.Assert(authors[1].Id, check.Equals, int64(2))
	c.Assert(authors[2].Id, check.Equals, int64(3))

	c.Assert(authors[0].Books, check.Equals, 2)
	c.Assert(authors[1].Books, check.Equals, 1)
	c.Assert(authors[2].Books, check.Equals, 1)
}

func (s *TestSuit) TestSearchWithAuthor(c *check.C) {
	authors := Search(map[string]interface{}{"author": 1})

	c.Assert(authors[0].Id, check.Equals, int64(1))
	c.Assert(authors[0].Books, check.Equals, 2)
}

func (s *TestSuit) TestSearchWithPublisher(c *check.C) {
	authors := Search(map[string]interface{}{"publisher": 2})

	c.Assert(authors[0].Id, check.Equals, int64(2))
	c.Assert(authors[0].Books, check.Equals, 1)
}

func (s *TestSuit) TestSearchWithCategory(c *check.C) {
	authors := Search(map[string]interface{}{"category": 1})

	c.Assert(authors[0].Id, check.Equals, int64(1))
	c.Assert(authors[1].Id, check.Equals, int64(3))

	c.Assert(authors[0].Books, check.Equals, 2)
	c.Assert(authors[1].Books, check.Equals, 1)
}

func (s *TestSuit) TestSearchWithSearch(c *check.C) {
	authors := Search(map[string]interface{}{"search": "The Hobbit"})

	c.Assert(authors[0].Id, check.Equals, int64(3))
	c.Assert(authors[0].Books, check.Equals, 1)
}

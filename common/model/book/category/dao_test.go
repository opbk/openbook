package category

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
	connection().Exec("ALTER SEQUENCE auto_id_categories RESTART WITH 1")
	connection().Exec("TRUNCATE categories")

	(&Category{0, 0, "", "Business", 0}).Save()
	(&Category{0, 0, "", "Fiction", 0}).Save()

	(&Category{0, 1, ">1", "Marketing", 0}).Save()
	(&Category{0, 3, ">1>3", "PR", 0}).Save()
	(&Category{0, 3, ">1>3", "IM", 0}).Save()

	(&Category{0, 2, ">2", "SF", 0}).Save()
	(&Category{0, 6, ">2>6", "Other SF", 0}).Save()
	(&Category{0, 7, ">2>6>7", "CyberPunk", 0}).Save()

	(&book.Book{0, "The Martian", 600, "en", "", "Six days ago, astronaut Mark Watney became one of the first people to walk on Mars.", "", "", time.Date(2013, time.March, 22, 0, 0, 0, 0, time.Local), time.Now(), 0, 1, []int64{1}, []int64{500}}).Save()
	(&book.Book{0, "The Egg", 500, "en", "", "You were on your way home when you died.It was a car accident.", "", "", time.Date(2014, time.November, 26, 0, 0, 0, 0, time.Local), time.Now(), 0, 1, []int64{1}, []int64{1}}).Save()
	(&book.Book{0, "Inside the Doomsday Machine", 400, "en", "", "Who understood the risk inherent in the assumption...", "", "", time.Date(2010, time.March, 15, 0, 0, 0, 0, time.Local), time.Now(), 0, 2, []int64{2}, []int64{2}}).Save()
	(&book.Book{0, "The Hobbit", 700, "en", "", "Like every other hobbit, Bilbo Baggins likes nothing...", "", "", time.Date(2012, time.September, 10, 0, 0, 0, 0, time.Local), time.Now(), 0, 1, []int64{1}, []int64{3}}).Save()
}

func (s *TestSuit) TearDownSuite(c *check.C) {
	connection().Exec("ALTER SEQUENCE auto_id_categories RESTART WITH 1")
	connection().Exec("TRUNCATE book_categories")
}

func (s *TestSuit) TestGetPath(c *check.C) {
	categories := List(8, 0)

	c.Check(len(categories[0].GetPath()), check.Equals, 0)
	c.Check(len(categories[1].GetPath()), check.Equals, 0)

	c.Check(categories[2].GetPath()[0].Id, check.Equals, int64(1))

	c.Check(categories[3].GetPath()[0].Id, check.Equals, int64(1))
	c.Check(categories[3].GetPath()[1].Id, check.Equals, int64(3))

	c.Check(categories[4].GetPath()[0].Id, check.Equals, int64(1))
	c.Check(categories[4].GetPath()[1].Id, check.Equals, int64(3))

	c.Check(categories[5].GetPath()[0].Id, check.Equals, int64(2))

	c.Check(categories[6].GetPath()[0].Id, check.Equals, int64(2))
	c.Check(categories[6].GetPath()[1].Id, check.Equals, int64(6))

	c.Check(categories[7].GetPath()[0].Id, check.Equals, int64(2))
	c.Check(categories[7].GetPath()[1].Id, check.Equals, int64(6))
	c.Check(categories[7].GetPath()[2].Id, check.Equals, int64(7))

}

func (s *TestSuit) TestGetChildCategories(c *check.C) {
	categories := List(8, 0)

	c.Check(categories[0].GetChildCategories()[0].Id, check.Equals, int64(3))

	c.Check(categories[1].GetChildCategories()[0].Id, check.Equals, int64(6))

	c.Check(categories[2].GetChildCategories()[0].Id, check.Equals, int64(5))
	c.Check(categories[2].GetChildCategories()[1].Id, check.Equals, int64(4))

	c.Check(categories[5].GetChildCategories()[0].Id, check.Equals, int64(7))

	c.Check(categories[6].GetChildCategories()[0].Id, check.Equals, int64(8))

}

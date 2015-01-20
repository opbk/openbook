package controller

import (
	"fmt"
	"path"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/check.v1"

	// "github.com/opbk/openbook/common/db"
	"github.com/opbk/openbook/common/configuration"
	"github.com/opbk/openbook/common/model/author"
	"github.com/opbk/openbook/common/model/book"
	"github.com/opbk/openbook/common/model/book/category"
)

type BookControllerTestSuit struct {
	FrontControllerTestSuit
}

var _ = check.Suite(new(BookControllerTestSuit))

func (s *BookControllerTestSuit) SetUpSuite(c *check.C) {
	config := configuration.LoadConfiguration("")
	config.Frontend.TemplatePath = path.Join("..", "..", config.Frontend.TemplatePath)

	router := mux.NewRouter()
	NewBookController().Routes(router)
	s.FrontControllerTestSuit.SetUpSuite(c, router)
}

func (s *BookControllerTestSuit) SetUpTest(c *check.C) {
	// (&author.Author{0, "Andy Weir"}).Save()
	// (&author.Author{0, "Michael Lewis"}).Save()
	// (&author.Author{0, "The Lord of the Rings"}).Save()
	s.FrontControllerTestSuit.SetUpTest(c)
}

func (s *BookControllerTestSuit) TearDownSuite(c *check.C) {
	s.FrontControllerTestSuit.TearDownSuite(c)
}

func (s *BookControllerTestSuit) TestSearch(c *check.C) {
	bookSearch = bookSearchMock
	authorListByBooks = authorListByBooksMock
	categoryListByParent = categoryListByParentMock
	categoryFind = categoryFindMock

	s.Get("/search")
}

func bookSearchMock(search map[string]interface{}, limit, offset int) []*book.Book {
	fmt.Println("Hello! I'm mock function")

	return []*book.Book{
		&book.Book{0, "The Martian", "Six days ago, astronaut Mark Watney became one of the first people to walk on Mars.", time.Date(2013, time.March, 22, 0, 0, 0, 0, time.Local), time.Now()},
		&book.Book{0, "The Egg", "You were on your way home when you died.It was a car accident.", time.Date(2014, time.November, 26, 0, 0, 0, 0, time.Local), time.Now()},
		&book.Book{0, "Inside the Doomsday Machine", "Who understood the risk inherent in the assumption...", time.Date(2010, time.March, 15, 0, 0, 0, 0, time.Local), time.Now()},
		&book.Book{0, "The Hobbit", "Like every other hobbit, Bilbo Baggins likes nothing...", time.Date(2012, time.September, 10, 0, 0, 0, 0, time.Local), time.Now()},
	}
}

func authorListByBooksMock(booksId []int64) []*author.Author {
	return make([]*author.Author, 0)
}

func categoryListByParentMock(categoryId int64) []*category.Category {
	return make([]*category.Category, 0)
}

func categoryFindMock(categoryId int64) *category.Category {
	return nil
}

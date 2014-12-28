package controller

import (
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/check.v1"

	"github.com/opbk/openbook/common/db"
	"github.com/opbk/openbook/common/model/book"
)

type BookControllerTestSuit struct {
	FrontControllerTestSuit
}

var _ = check.Suite(new(BookControllerTestSuit))

func (s *BookControllerTestSuit) SetUpSuite(c *check.C) {
	router := mux.NewRouter()
	NewBookController().Routes(router)
	s.FrontControllerTestSuit.SetUpSuite(c, router)
}

func (s *BookControllerTestSuit) SetUpTest(c *check.C) {
	db.Connection().Exec("ALTER SEQUENCE auto_id_books RESTART WITH 1")
	db.Connection().Exec("TRUNCATE books")
	db.Connection().Exec("TRUNCATE book_categories")

	(&book.Book{0, "The Martian", "Six days ago, astronaut Mark Watney became one of the first people to walk on Mars.", time.Date(2013, time.March, 22, 0, 0, 0, 0, time.Local)}).Save()
	(&book.Book{0, "The Egg", "You were on your way home when you died.It was a car accident.", time.Date(2014, time.November, 26, 0, 0, 0, 0, time.Local)}).Save()
	(&book.Book{0, "Inside the Doomsday Machine", "Who understood the risk inherent in the assumption...", time.Date(2010, time.March, 15, 0, 0, 0, 0, time.Local)}).Save()
	(&book.Book{0, "The Hobbit", "Like every other hobbit, Bilbo Baggins likes nothing...", time.Date(2012, time.September, 10, 0, 0, 0, 0, time.Local)}).Save()

	book.AddBookToCategory(1, 1)
	book.AddBookToCategory(2, 1)
	book.AddBookToCategory(4, 1)

	s.FrontControllerTestSuit.SetUpTest(c)
}

func (s *BookControllerTestSuit) TearDownSuite(c *check.C) {
	db.Connection().Exec("ALTER SEQUENCE auto_id_books RESTART WITH 1")
	db.Connection().Exec("TRUNCATE books")
	db.Connection().Exec("TRUNCATE book_categories")
	s.FrontControllerTestSuit.TearDownSuite(c)
}

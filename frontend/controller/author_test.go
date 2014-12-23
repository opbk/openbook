package controller

import (
	"github.com/gorilla/mux"
	"gopkg.in/check.v1"

	"github.com/opbk/openbook/common/db"
	"github.com/opbk/openbook/common/model/author"
)

type AuthorControllerTestSuit struct {
	FrontControllerTestSuit
}

var _ = check.Suite(new(AuthorControllerTestSuit))

func (s *AuthorControllerTestSuit) SetUpSuite(c *check.C) {
	router := mux.NewRouter()
	NewAuthorController().Routes(router)
	s.FrontControllerTestSuit.SetUpSuite(c, router)
}

func (s *AuthorControllerTestSuit) SetUpTest(c *check.C) {
	db.Connection().Exec("ALTER SEQUENCE auto_id_authors RESTART WITH 1")
	db.Connection().Exec("TRUNCATE authors")

	(&author.Author{0, "Suzanne Collins"}).Save()
	(&author.Author{0, "Andy Weir"}).Save()
	(&author.Author{0, "George R. R. Martin"}).Save()
	s.FrontControllerTestSuit.SetUpTest(c)
}

func (s *AuthorControllerTestSuit) TearDownSuite(c *check.C) {
	db.Connection().Exec("ALTER SEQUENCE auto_id_authors RESTART WITH 1")
	db.Connection().Exec("TRUNCATE authors")
	s.FrontControllerTestSuit.TearDownSuite(c)
}

func (s *AuthorControllerTestSuit) TestList(c *check.C) {
	expected := "[" +
		"{\"id\":1,\"name\":\"Suzanne Collins\"}," +
		"{\"id\":2,\"name\":\"Andy Weir\"}," +
		"{\"id\":3,\"name\":\"George R. R. Martin\"}]"

	res, _ := s.Get("/authors")
	c.Assert(res, check.Equals, expected)
}

func (s *AuthorControllerTestSuit) TestGet(c *check.C) {
	expected := "{\"id\":1,\"name\":\"Suzanne Collins\"}"
	res, _ := s.Get("/authors/1")
	c.Assert(res, check.Equals, expected)
}

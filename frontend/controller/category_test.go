package controller

import (
	"github.com/gorilla/mux"
	"gopkg.in/check.v1"

	"github.com/opbk/openbook/common/db"
	"github.com/opbk/openbook/common/model/book/category"
)

type CategoryControllerTestSuit struct {
	FrontControllerTestSuit
}

var _ = check.Suite(new(CategoryControllerTestSuit))

func (s *CategoryControllerTestSuit) SetUpSuite(c *check.C) {
	router := mux.NewRouter()
	NewCategoryController().Routes(router)
	s.FrontControllerTestSuit.SetUpSuite(c, router)
}

func (s *CategoryControllerTestSuit) SetUpTest(c *check.C) {
	db.Connection().Exec("ALTER SEQUENCE auto_id_categories RESTART WITH 1")
	db.Connection().Exec("TRUNCATE categories")

	(&category.Category{0, 0, "Sci-Fi & Fantasy"}).Save()
	(&category.Category{0, 0, "History"}).Save()
	(&category.Category{0, 2, "20th Century"}).Save()
	(&category.Category{0, 2, "21st Century"}).Save()
	s.FrontControllerTestSuit.SetUpTest(c)
}

func (s *CategoryControllerTestSuit) TearDownSuite(c *check.C) {
	db.Connection().Exec("ALTER SEQUENCE auto_id_categories RESTART WITH 1")
	db.Connection().Exec("TRUNCATE categories")
	s.FrontControllerTestSuit.TearDownSuite(c)
}

func (s *CategoryControllerTestSuit) TestListRootCategory(c *check.C) {
	expected := "[" +
		"{\"id\":1,\"category_id\":0,\"name\":\"Sci-Fi & Fantasy\"}," +
		"{\"id\":2,\"category_id\":0,\"name\":\"History\"}]"

	res, _ := s.Get("/categories")
	c.Assert(res, check.Equals, expected)
}

func (s *CategoryControllerTestSuit) TestListChildCategory(c *check.C) {
	expected := "[" +
		"{\"id\":2,\"category_id\":2,\"name\":\"20th Century\"}," +
		"{\"id\":3,\"category_id\":2,\"name\":\"21st Century\"}]"

	res, _ := s.Get("/categories/2")
	c.Assert(res, check.Equals, expected)
}

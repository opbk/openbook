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
	s.FrontControllerTestSuit.SetUpTest(c)
}

func (s *BookControllerTestSuit) TearDownSuite(c *check.C) {
	s.FrontControllerTestSuit.TearDownSuite(c)
}

func (s *BookControllerTestSuit) TestSearch(c *check.C) {

}

func bookSearchMock(search map[string]interface{}, limit, offset int) []*book.Book {
	return make([]book.Book)
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

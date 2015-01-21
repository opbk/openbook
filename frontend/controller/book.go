package controller

import (
	"html/template"
	"math"
	"net/http"
	"path"
	"time"

	// logger "github.com/cihub/seelog"
	"github.com/gorilla/mux"

	"github.com/opbk/openbook/common/configuration"
	"github.com/opbk/openbook/common/model/author"
	"github.com/opbk/openbook/common/model/book"
	"github.com/opbk/openbook/common/model/book/category"
	"github.com/opbk/openbook/common/model/publisher"
	"github.com/opbk/openbook/common/web"
)

type BookController struct {
	FrontendController
	DefaultLimitPerPage int
}

var tfns = template.FuncMap{
	"add":    func(a, b int) int { return a + b },
	"sub":    func(a, b int) int { return a - b },
	"xrange": func(a int) []int { return make([]int, a) },
	"pagination": func(total, limit int) []map[string]int {
		total = int(math.Ceil(float64(total) / float64(limit)))
		result := make([]map[string]int, 0)
		for i := 0; i < total; i++ {
			result = append(result, map[string]int{
				"page":   i + 1,
				"offset": i * limit,
			})
		}

		return result
	},
	"html": func(a string) template.HTML { return template.HTML(a) },
}

func NewBookController() *BookController {
	c := new(BookController)
	c.DefaultLimitPerPage = 10

	tPath := configuration.GetConfig().Frontend.TemplatePath
	c.template = template.Must(template.New("index").Funcs(tfns).Delims("{%", "%}").ParseFiles(
		path.Join(tPath, "header.html"),
		path.Join(tPath, "footer.html"),
		path.Join(tPath, "book", "search.html"),
		path.Join(tPath, "book", "book.html"),
	))

	return c
}

func (c *BookController) Routes(router *mux.Router) {
	router.HandleFunc("/search", c.Search)
	router.HandleFunc("/book/{id:[0-9]+}", c.Book)
}

func (c *BookController) Book(rw http.ResponseWriter, req *http.Request) {
	request := web.NewRequest(req)
	b := bookFind(request.GetInt64("id"))
	if b == nil {
		http.NotFound(rw, req)
	}

	prices := b.Prices()
	authors := authorMapById(b.AuthorsId)
	publishers := publisherMapById([]int64{b.PublisherId})

	var path []*category.Category
	if c := categoryFind(request.GetInt64("c")); c != nil {
		path = c.GetPath()
		path = append(path, c)
	}

	dueDate := time.Now().AddDate(0, 1, 0)

	c.template.ExecuteTemplate(rw, "book", map[string]interface{}{
		"book":          b,
		"prices":        prices,
		"authors":       authors,
		"publishers":    publishers,
		"path":          path,
		"recomendation": make([]*book.Book, 0),
		"dueDate":       dueDate,
	})
}

func (c *BookController) Search(rw http.ResponseWriter, req *http.Request) {
	request := web.NewRequest(req)

	limit := request.GetInt("l", c.DefaultLimitPerPage)
	offset := request.GetInt("f")

	var path []*category.Category
	if c := categoryFind(request.GetInt64("c")); c != nil {
		path = c.GetPath()
		path = append(path, c)
	}

	filter := buildSearchMap(request)
	books := bookSearch(filter, limit, offset)
	booksCount := bookSearchCount(filter)

	searchAuthors := authorSearch(filter)
	searchPublisher := publisherSearch(filter)
	searchCategories := categorySearch(filter)

	authors := make(map[int64]*author.Author)
	publishers := make(map[int64]*publisher.Publisher)
	if len(books) > 0 {
		authorsId := make([]int64, len(books))
		publishersId := make([]int64, len(books))
		for _, book := range books {
			authorsId = append(authorsId, book.AuthorsId...)
			publishersId = append(publishersId, book.PublisherId)
		}

		authors = authorMapById(authorsId)
		publishers = publisherMapById(publishersId)
	}

	c.template.ExecuteTemplate(rw, "search", map[string]interface{}{
		"books": map[string]interface{}{
			"books":      books,
			"authors":    authors,
			"publishers": publishers,
		},
		"author":     request.GetInt64("a"),
		"publisher":  request.GetInt64("p"),
		"category":   request.GetInt64("c"),
		"search":     request.GetString("s"),
		"authors":    searchAuthors,
		"publishers": searchPublisher,
		"categories": searchCategories,
		"path":       path,
		"pagination": map[string]int{
			"total":  booksCount,
			"limit":  limit,
			"offset": offset,
		},
	})
}

func buildSearchMap(req *web.Request) map[string]interface{} {
	filter := make(map[string]interface{})
	if search := req.GetString("s"); search != "" {
		filter["search"] = search
	}
	if categoryId := req.GetInt64("c"); categoryId != 0 {
		filter["category"] = categoryId
	}
	if authorId := req.GetInt64("a"); authorId != 0 {
		filter["author"] = authorId
	}
	if publisherId := req.GetInt64("p"); publisherId != 0 {
		filter["publisher"] = publisherId
	}
	if release := req.GetString("r"); release != "" {
		filter["release"] = release
	}

	return filter
}

var bookSearch = func(filter map[string]interface{}, limit, offset int) []*book.Book {
	return book.Search(filter, limit, offset)
}

var bookSearchCount = func(filter map[string]interface{}) int {
	return book.SearchCount(filter)
}

var bookFind = func(id int64) *book.Book {
	return book.Find(id)
}

var authorMapById = func(ids []int64) map[int64]*author.Author {
	return author.MapById(ids)
}

var authorFind = func(id int64) *author.Author {
	return author.Find(id)
}

var authorSearch = func(filter map[string]interface{}) []*author.Author {
	return author.Search(filter)
}

var publisherMapById = func(ids []int64) map[int64]*publisher.Publisher {
	return publisher.MapById(ids)
}

var publisherSearch = func(filter map[string]interface{}) []*publisher.Publisher {
	return publisher.Search(filter)
}

var categoryListChildCategories = func(id int64) []*category.Category {
	return category.ListChildCategories(id)
}

var categoryFind = func(id int64) *category.Category {
	return category.Find(id)
}

var categorySearch = func(filter map[string]interface{}) []*category.Category {
	return category.Search(filter)
}

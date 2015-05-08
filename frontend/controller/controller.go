package controller

import (
	"html/template"
	"math"
	"net/http"
	"path"

	"github.com/opbk/openbook/common/configuration"
	"github.com/opbk/openbook/common/model/user"
	"github.com/opbk/openbook/common/web"
	"github.com/opbk/openbook/common/web/auth"
)

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

type FrontendController struct {
	template *template.Template
}

func (c *FrontendController) Template(tpl ...*template.Template) *template.Template {
	if len(tpl) > 0 {
		c.template = tpl[0]
	}

	return c.template
}

func (c *FrontendController) ExecuteTemplate(rw http.ResponseWriter, req *http.Request, name string, params map[string]interface{}) {
	user := c.getUser(req)
	params["user"] = user
	if user != nil {
		params["subscription"] = user.Subscription()
	}

	params["location"] = req.URL.String()
	c.Template().ExecuteTemplate(rw, name, params)
}

func (c *FrontendController) getUser(req *http.Request) *user.User {
	return auth.Get(req)
}

func GetTemplate() *template.Template {
	tPath := configuration.GetConfig().Frontend.TemplatePath
	return template.Must(template.New("index").Funcs(tfns).Delims("{%", "%}").ParseFiles(
		path.Join(tPath, "header.html"),
		path.Join(tPath, "footer.html"),
		path.Join(tPath, "howitworks.html"),
		path.Join(tPath, "order.html"),
		path.Join(tPath, "sign", "signup.html"),
		path.Join(tPath, "book", "search.html"),
		path.Join(tPath, "book", "book.html"),
		path.Join(tPath, "user", "history.html"),
		path.Join(tPath, "user", "wishlist.html"),
		path.Join(tPath, "user", "subscribe.html"),
		path.Join(tPath, "email", "signup.html"),
		path.Join(tPath, "email", "user_order.html"),
		path.Join(tPath, "email", "admin_order.html"),
	))
}

func GetControllers() []web.Controller {
	controllers := []web.Controller{
		NewIndexController(),
		NewSignController(),
		NewBookController(),
		NewOrderController(),
		NewUserController(),
		NewAddressController(),
		NewSubscriptionController(),
	}

	return controllers
}

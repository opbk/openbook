package controller

import (
	"encoding/json"
	"html/template"
	"math"
	"net/http"
	"path"

	"github.com/opbk/openbook/common/configuration"
	"github.com/opbk/openbook/common/model/order"
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
	"set": func(m map[string]interface{}, key string, value interface{}) string {
		m[key] = value
		return ""
	},
	"json": func(o interface{}) string {
		bytes, _ := json.Marshal(o)
		return string(bytes)
	},
	"ending": func(count int, options ...string) string {
		index := count % 100
		if index >= 11 && index <= 14 {
			index = 0
		} else {
			index = index % 10
			if index < 5 {
				if index > 2 {
					index = 2
				}
			} else {
				index = 0
			}
		}

		return options[index]
	},
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
		params["ordersOnhand"] = order.CountByUserAndStatus(user.Id, order.ONHAND)
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
		path.Join(tPath, "about.html"),
		path.Join(tPath, "sign", "signup.html"),
		path.Join(tPath, "user", "wishlist.html"),
		path.Join(tPath, "book", "search.html"),
		path.Join(tPath, "book", "book.html"),
		path.Join(tPath, "order", "order.html"),
		path.Join(tPath, "order", "history.html"),
		path.Join(tPath, "subscription", "subscribe.html"),
		path.Join(tPath, "email", "signup.html"),
		path.Join(tPath, "email", "order_user.html"),
		path.Join(tPath, "email", "order_admin.html"),
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

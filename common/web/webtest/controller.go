package webtest

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gorilla/mux"
	"gopkg.in/check.v1"

	// "github.com/opbk/openbook/common/model/user"
	// "github.com/opbk/openbook/common/web/auth"
)

type ControllerTestSuit struct {
	server *httptest.Server
	//authWrapper *AuthWrapper
	c *check.C
}

func (s *ControllerTestSuit) SetUpSuite(c *check.C, router *mux.Router) {
	//s.authWrapper = NewAuthWrapper(router)
	s.server = httptest.NewServer(router)
	s.c = c
}

func (s *ControllerTestSuit) TearDownSuite(c *check.C) {
	s.server.Close()
}

// func (s *ControllerTestSuit) WithUser(u *user.User) {
// 	s.authWrapper.WithUser(u)
// }

func (s *ControllerTestSuit) Get(url string) (string, *http.Response) {
	res, err := http.Get(s.server.URL + url)
	if err != nil && !strings.Contains(err.Error(), "not following redirect") {
		s.c.Fatalf("Error while get request: %s", err)
	}

	result, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	return string(result), res
}

func (s *ControllerTestSuit) Post(url string, body string) (string, *http.Response) {
	res, err := http.Post(s.server.URL+url, "application/json", strings.NewReader(body))
	if err != nil {
		s.c.Fatalf("Error while post request: %s", err)
	}

	result, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	return string(result), res
}

func (s *ControllerTestSuit) Put(url string, body string) (string, *http.Response) {
	req, _ := http.NewRequest("PUT", s.server.URL+url, strings.NewReader(body))
	res, err := (&http.Client{}).Do(req)
	if err != nil {
		s.c.Fatalf("Error while put request: %s", err)
	}

	result, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	return string(result), res
}

func (s *ControllerTestSuit) Delete(url string) (string, *http.Response) {
	req, _ := http.NewRequest("DELETE", s.server.URL+url, nil)
	res, err := (&http.Client{}).Do(req)
	if err != nil {
		s.c.Fatalf("Error while put request: %s", err)
	}

	result, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	return string(result), res
}

// type AuthWrapper struct {
// 	router *mux.Router
// 	user   *user.User
// }

// func NewAuthWrapper(router *mux.Router) *AuthWrapper {
// 	return &AuthWrapper{router: router}
// }

// func (w *AuthWrapper) WithUser(u *user.User) {
// 	w.user = u
// }

// func (w *AuthWrapper) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
// 	if w.user != nil {
// 		auth.Set(w.user, rw, req)
// 	}

// 	w.router.ServeHTTP(rw, req)
// 	w.user = nil
// }

package webtest

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gorilla/mux"
	"gopkg.in/check.v1"

	"github.com/opbk/openbook/common/configuration"
	"github.com/opbk/openbook/common/db"
)

type ControllerTestSuit struct {
	server *httptest.Server
	c      *check.C
	router *mux.Router
}

func (s *ControllerTestSuit) GetRouter() *mux.Router {
	if s.router == nil {
		s.router = mux.NewRouter()
	}

	return s.router
}

func (s *ControllerTestSuit) SetUpSuite(c *check.C) {
	config := configuration.LoadConfiguration("")
	db.InitDbConnection(config.Db.Driver, config.Db.Connection)

	s.server = httptest.NewServer(s.router)
	s.c = c
}

func (s *ControllerTestSuit) TearDownSuite(c *check.C) {
	s.server.Close()
}

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

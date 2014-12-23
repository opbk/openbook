package controller

import (
	"testing"

	"github.com/gorilla/mux"
	"gopkg.in/check.v1"

	"github.com/opbk/openbook/common/configuration"
	"github.com/opbk/openbook/common/db"
	"github.com/opbk/openbook/common/web/webtest"
)

func Test(t *testing.T) { check.TestingT(t) }

type FrontControllerTestSuit struct {
	webtest.ControllerTestSuit
}

func (s *FrontControllerTestSuit) SetUpSuite(c *check.C, router *mux.Router) {
	config := configuration.LoadConfiguration("")
	db.InitDbConnection(config.Db.Driver, config.Db.Connection)
	s.ControllerTestSuit.SetUpSuite(c, router)
}

func (s *FrontControllerTestSuit) SetUpTest(c *check.C) {

}

func (s *FrontControllerTestSuit) TearDownSuite(c *check.C) {
	s.ControllerTestSuit.TearDownSuite(c)
}

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

type FrontendControllerTestSuit struct {
	webtest.ControllerTestSuit
}

func (s *FrontendControllerTestSuit) SetUpSuite(c *check.C, router *mux.Router) {
	config := configuration.GetConfig()
	if config == nil {
		config = configuration.LoadConfiguration("")
	}

	db.InitDbConnection(config.Db.Driver, config.Db.Connection)
	s.ControllerTestSuit.SetUpSuite(c, router)
}

func (s *FrontendControllerTestSuit) SetUpTest(c *check.C) {

}

func (s *FrontendControllerTestSuit) TearDownSuite(c *check.C) {
	s.ControllerTestSuit.TearDownSuite(c)
}

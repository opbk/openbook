package payment

import (
	"gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) {
	check.Suite(new(YandexControllerTestSuite))
	check.TestingT(t)
}

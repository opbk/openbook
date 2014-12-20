package configuration

import (
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	config := LoadConfiguration("")
	if config.Db.Connection != "postgres://developer:developer@localhost/openbook" ||
		config.Db.Driver != "postgres" {
		t.Fail()
	}
}

func TestLoadConfigFile(t *testing.T) {
	config := LoadConfiguration("config.gcfg")
	if config.Db.Connection != "postgres://developer:developer@localhost/openbook" ||
		config.Db.Driver != "postgres" {
		t.Fail()
	}
}

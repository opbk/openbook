package configuration

import (
	"code.google.com/p/gcfg"
	logging "github.com/cihub/seelog"
)

type Frontend struct {
	HttpPort     string
	TemplatePath string
}

type Backend struct {
	HttpPort     string
	TemplatePath string
}

type Db struct {
	Driver     string
	Connection string
}

type Main struct {
	SendDelay string
	LogFile   string
	MaxProc   int
}

type Config struct {
	Frontend
	Backend
	Db
	Main
}

const defaultConfig = `
[db]
driver     = postgres
connection = postgres://developer:developer@localhost/newsgun

[frontend]
httpport       = 8089
templatepath   = resources/frontend/templates/

[backend]
httpport = 8090
templatepath   = resources/backend/templates/

[main]
senddelay = 0.1s
logfile = seelog.xml
maxproc = 8
`

var cfg Config

func LoadConfiguration(file string) *Config {
	var err error

	if file != "" {
		err = gcfg.ReadFileInto(&cfg, file)
	} else {
		err = gcfg.ReadStringInto(&cfg, defaultConfig)
	}

	if err != nil {
		logging.Criticalf("Error while reading configuration file: %s", err)
	}

	return &cfg
}

func GetConfig() *Config {
	return &cfg
}

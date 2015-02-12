package configuration

import (
	"code.google.com/p/gcfg"
	logging "github.com/cihub/seelog"
)

type Frontend struct {
	HttpPort     string
	TemplatePath string
	StaticPath   string
	UploadPath   string
}

type Backend struct {
	HttpPort     string
	TemplatePath string
}

type Db struct {
	Driver     string
	Connection string
}

type Aws struct {
	Region    string
	AccessKey string
	SecretKey string
	Bucket    string
}

type EmailSender struct {
	From string
}

type Main struct {
	LogFile string
	MaxProc int
	Domain  string
}

type Config struct {
	Frontend
	Backend
	Db
	Aws
	EmailSender
	Main
}

const defaultConfig = `
[db]
driver     = postgres
connection = postgres://developer:developer@localhost/openbook

[frontend]
httpport     = 8089
templatepath = resources/frontend/templates/
staticpath   = resources/frontend/static/
uploadpath   = resources/frontend/templates/

[backend]
httpport     = 8090
templatepath = resources/backend/templates/

[aws]
region = us-east-1
accesskey = 
secretkey = 

[emailsender]
from = noreply@opbook.rog

[main]
logfile = seelog.xml
maxproc = 8
domain = opbook.org
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

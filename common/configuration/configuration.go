package configuration

import (
	"flag"
	"fmt"

	"code.google.com/p/gcfg"
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

var config Config
var TEST_CONFIG_FAILE string

func LoadConfiguration(file string) *Config {
	var err error
	if file == "" {
		file = TEST_CONFIG_FAILE
		if file == "" {
			panic("flag -test.config is required")
		}
	}

	err = gcfg.ReadFileInto(&config, file)
	if err != nil {
		panic(fmt.Sprintf("Error while reading configuration file: %s", err))
	}

	return &config
}

func GetConfig() *Config {
	return &config
}

func init() {
	flag.StringVar(&TEST_CONFIG_FAILE, "test.config", "", "test configuration file")
}

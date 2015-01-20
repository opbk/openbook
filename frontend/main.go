package main

import (
	"flag"
	"fmt"
	"net/http"
	"runtime"

	logger "github.com/cihub/seelog"
	"github.com/gorilla/mux"

	"github.com/opbk/openbook/common/configuration"
	"github.com/opbk/openbook/common/db"
	"github.com/opbk/openbook/frontend/controller"
)

func initLogging(config *configuration.Config) {
	newLogger, _ := logger.LoggerFromConfigAsFile(config.Main.LogFile)
	logger.ReplaceLogger(newLogger)
}

func initDataBases(config *configuration.Config) {
	db.InitDbConnection(config.Db.Driver, config.Db.Connection)
}

func initControllers(router *mux.Router) {
	for _, controller := range controller.GetControllers() {
		controller.Routes(router)
	}
}

func initStaticHandler(router *mux.Router, config *configuration.Config) {
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(config.Frontend.StaticPath))))
	router.PathPrefix("/upload/").Handler(http.StripPrefix("/upload/", http.FileServer(http.Dir(config.Frontend.UploadPath))))
}

func main() {
	var configFile *string = flag.String("config", "/etc/openbook/config.gcfg", "configuration file")
	flag.Parse()

	config := configuration.LoadConfiguration(*configFile)
	runtime.GOMAXPROCS(config.Main.MaxProc)
	initLogging(config)
	initDataBases(config)

	router := mux.NewRouter()
	initStaticHandler(router, config)
	initControllers(router)

	logger.Infof("Listening webserver on port %s", config.Frontend.HttpPort)
	err := http.ListenAndServe(fmt.Sprintf(":%s", config.Frontend.HttpPort), router)
	if err != nil {
		panic(err)
	}
}

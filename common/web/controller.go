package web

import "github.com/gorilla/mux"

type Controller interface {
	Routes(router *mux.Router)
}

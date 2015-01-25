package web

import (
	"github.com/gorilla/mux"
	"html/template"
)

type Controller interface {
	Routes(router *mux.Router)
	Template(tpl ...*template.Template) *template.Template
}

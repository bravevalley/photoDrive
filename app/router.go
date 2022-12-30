package app

import (
	"html/template"

	"github.com/gorilla/mux"
)

type Router struct {
	Mux *mux.Router
}

func (router *Router) Init(tpl *template.Template) {
	router.Mux = mux.NewRouter()
	router.ServeRoute(tpl)
}

func (router *Router) ServeRoute(tpl *template.Template) {
	handler := &Multiplexer{
		Template: tpl,
	}
	router.Mux.HandleFunc("/", handler.IndexHandler)
}
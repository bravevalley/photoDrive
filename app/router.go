package app

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

// Create a Router struct to set up the router
type Router struct {
	Mux *mux.Router
}

// Init initializes the router for traffic by assigning the multiplexer
func (router *Router) Init(tpl *template.Template) {
	router.Mux = mux.NewRouter()
	router.ServeRoute(tpl)
}

// ServeRoute route traffic and serve content from the handler
func (router *Router) ServeRoute(tpl *template.Template) {
	handler := &Multiplexer{
		Template: tpl,
	}
	router.Mux.HandleFunc("/", handler.IndexHandler)
	router.Mux.PathPrefix("/assets/").Handler (http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

}
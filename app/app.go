package app

import (
	"log"
	"net/http"
	"html/template"
)

// Create an app struct that serves as api to the main func
type App struct {
	Router *Router
	Server *http.Server
}

// Init initiates the App by creating the router struct
func (app *App) Init(tpl *template.Template)  {
	app.Router = &Router{}
	app.Router.Init(tpl)
	
}

// Run creates the server configuration and starts the server
func (app *App) Run()  {
	app.Server = &http.Server{
		Addr:    ":8080",
		Handler: app.Router.Mux,
	}

	log.Fatal(app.Server.ListenAndServe())
}
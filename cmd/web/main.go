package main

import (
	"github.com/EraldBa/webApp/pkg/config"
	"github.com/EraldBa/webApp/pkg/handlers"
	"github.com/EraldBa/webApp/pkg/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

const portNumber string = ":8000"

var app config.AppConfig

func main() {
	// Creating app config in main

	var err error
	// Set to true if in production, for now it's false
	app.InProduction = false

	// Initializing session
	app.Session = scs.New()
	// Setting session lifetime
	app.Session.Lifetime = 1 * time.Hour
	// Session ends when user closes browser, by setting Cookie.Persist to false
	app.Session.Cookie.Persist = false
	//
	app.Session.Cookie.SameSite = http.SameSiteLaxMode
	// Set to true if using https and need to encrypt cookies sent, for
	// now it's set to false
	app.Session.Cookie.Secure = app.InProduction

	app.TemplateCache, err = render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cant create template cache", err)
	}

	app.UseCache = true

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// Passing app config to render package
	render.NewTemplates(&app)

	srv := http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	log.Println("Listening on port", portNumber)

	err = srv.ListenAndServe()

	log.Fatal(err)

}

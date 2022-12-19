package main

import (
	"github.com/EraldBa/webApp/pkg/config"
	"github.com/EraldBa/webApp/pkg/driver"
	"github.com/EraldBa/webApp/pkg/handlers"
	"github.com/EraldBa/webApp/pkg/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

const (
	dbInfo     = "host=localhost post=5432 dbname=fitbuddy user=erald password=95486672"
	portNumber = ":8080"
)

var app config.AppConfig

func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	srv := http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	log.Println("Listening on port", portNumber)

	err = srv.ListenAndServe()

	log.Fatal(err)

}

func run() (*driver.DB, error) {
	var err error
	// Set to true if in production, for now it's false
	app.InProduction = false

	// Initializing session
	app.Session = scs.New()
	// Setting session lifetime
	app.Session.Lifetime = 1 * time.Hour
	// Session ends when user closes browser, by setting Cookie.Persist to false
	app.Session.Cookie.Persist = true

	app.Session.Cookie.SameSite = http.SameSiteLaxMode
	// Set to true if using https and need to encrypt cookies sent, for
	// now it's set to false
	app.Session.Cookie.Secure = app.InProduction

	log.Println("Connecting to database...")
	db, err := driver.ConnectDB(dbInfo)
	if err != nil {
		log.Fatal("Cannot connect to database. Program is exiting...")
		return nil, err
	}
	log.Println("Connected to database!")

	app.TemplateCache, err = render.CreateTemplateCache()
	if err != nil {
		return nil, err
	}

	app.UseCache = false

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)

	// Passing app config to render package
	render.NewTemplates(&app)
	return db, nil
}

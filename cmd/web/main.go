package main

import (
	"flag"
	"fmt"
	"github.com/EraldBa/webApp/pkg/config"
	"github.com/EraldBa/webApp/pkg/driver"
	"github.com/EraldBa/webApp/pkg/handlers"
	"github.com/EraldBa/webApp/pkg/helpers"
	"github.com/EraldBa/webApp/pkg/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"os"
	"time"
)

const portNumber = ":8080"

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

	dbHost := flag.String("dbhost", "localhost", "Database host")
	dbPort := flag.String("dbport", "5432", "Database port number")
	dbName := flag.String("dbname", "", "Database name")
	dbUser := flag.String("dbuser", "", "Database user")
	dbPassword := flag.String("dbpassword", "", "Database password")
	dbSSL := flag.String("dbssl", "disable", "Database SSL settings (disable, prefer, require)")

	inProduction := flag.Bool("production", true, "Application is in production")
	useCache := flag.Bool("cache", true, "Use template cache")
	flag.Parse()

	if *dbName == "" || *dbUser == "" {
		log.Fatal("Missing required fields: dbname and/or dbuser.")
	}
	// dsn is the database connection info\
	dsn := "host=%s port=%s dbname=%s user=%s password=%s sslmode=%s"
	dsn = fmt.Sprintf(dsn, *dbHost, *dbPort, *dbName, *dbUser, *dbPassword, *dbSSL)
	// Set to true if in production, for now it's false
	app.InProduction = *inProduction

	app.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	app.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

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

	app.TemplateCache, err = render.CreateTemplateCache()
	if err != nil {
		return nil, err
	}
	app.UseCache = *useCache

	log.Println("Connecting to database...")
	db := driver.ConnectDB(dsn)
	log.Println("Connected to database!")

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	helpers.NewHelpers(&app)
	// Passing app config to render package
	render.NewRenderer(&app)
	return db, nil
}

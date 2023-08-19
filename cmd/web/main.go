package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"

	"github.com/dzonib/golang-app-with-templates/pkg/config"
	"github.com/dzonib/golang-app-with-templates/pkg/handlers"
	"github.com/dzonib/golang-app-with-templates/pkg/templates"
)

const portNumber = ":8080"

var app config.AppConfig

var sessionManager *scs.SessionManager

func main() {
	// change this to true when in production
	app.InProduction = false

	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = false // we want this true in prod

	app.Session = sessionManager

	tc, err := templates.CreateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	// we don't need to cache in dev, we want latest page to show always
	app.UseCache = false

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	templates.InitiatetemplateConfig(&app)

	//	http.HandleFunc("/", repo.Home)
	//	http.HandleFunc("/about", repo.About)

	fmt.Println(fmt.Sprintf("starting application on port: %s", portNumber))

	//	err = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: Routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

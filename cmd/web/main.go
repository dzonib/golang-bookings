package main

import (
	"fmt"
	"log"
	"net/http"


	"github.com/dzonib/golang-app-with-templates/pkg/config"
	"github.com/dzonib/golang-app-with-templates/pkg/handlers"
	"github.com/dzonib/golang-app-with-templates/pkg/templates"
)

const portNumber = ":8080"

func main() {

	var app config.AppConfig

	tc, err := templates.CreateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	// we dont need cache in dev, we want latest page to show always
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

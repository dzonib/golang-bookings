package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dzonib/golang-app-with-templates/pkg/render"

	"github.com/dzonib/golang-app-with-templates/pkg/config"

	"github.com/dzonib/golang-app-with-templates/pkg/handlers"
)

const portNumber = ":8080"

func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	// we dont need cache in dev, we want latest page to show always
	app.UseCache = false

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	render.InitiateRenderConfig(&app)

	http.HandleFunc("/", repo.Home)
	http.HandleFunc("/about", repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port: %s", portNumber))

	err = http.ListenAndServe(portNumber, nil)
	if err != nil {
		panic(err)
	}
}

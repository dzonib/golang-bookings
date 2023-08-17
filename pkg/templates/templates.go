package templates

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/dzonib/golang-app-with-templates/pkg/models"

	"github.com/dzonib/golang-app-with-templates/pkg/config"
)

var app *config.AppConfig

// InitiatetemplateConfig sets a config for a template package
func InitiatetemplateConfig(a *config.AppConfig) {
	app = a
}

// templateTemplateTest does not return anything, it writes everything in response writer
// version without caching
//func templateTemplateTest(w http.ResponseWriter, tmpl string) {
//	// this can be expensive, we are solvingthis issue bellow
//	parsedTemplate, err := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
//
//	if err != nil {
//		http.Error(w, "Error parsing template: "+err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	err = parsedTemplate.Execute(w, nil)
//
//	if err != nil {
//		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
//	}
//}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate looking teplates, layouts, partials, and make them automatically populate the template cache (more complexed version of caching)
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template
	var err error
	// get the template cache from the app config

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, err = CreateCache()

		if err != nil {
			log.Println(err)
		}
	}

	// create template cache
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	// get requested template from cache
	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	// arbitrary step for better error checking
	buf := new(bytes.Buffer)

	//useful when we add more logic for passing td to specific templates
	td = AddDefaultData(td)

	// try to execute and check if it works
	err = t.Execute(buf, td)

	if err != nil {
		// tells us that error is comming from the map
		log.Println(err)
	}
	// template the template

	_, err = buf.WriteTo(w)

	if err != nil {
		log.Println(err)
	}

}

func CreateCache() (map[string]*template.Template, error) {
	//	myCache := make(map[string]*template.RenderTemplate)
	// same as using make keyword
	myCache := map[string]*template.Template{}

	// get all of the files named  *.page.tmpl from ./templates
	// we need to parse them first
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	//range through all files ending with *.page.tmpl

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}

// template cache
//var tc = make(map[string]*template.RenderTemplate)
//
//// templateTemplateCached version with caching, templates templates and caches them
//func templateTemplateCached(w http.ResponseWriter, t string) {
//	var tmpl *template.RenderTemplate
//	var err error
//
//	// check if we already have the template in cache
//	_, inMap := tc[t]
//	if !inMap {
//		log.Println("Creating teml and adding to cache")
//		// need to create a template
//		err = createTemplateCacheOld(t)
//		if err != nil {
//			log.Println(err)
//		}
//	} else {
//		// we have the template in the cache
//		log.Println("Using cached template")
//	}
//
//	tmpl = tc[t]
//
//	err = tmpl.Execute(w, nil)
//	if err != nil {
//		log.Println(err)
//	}
//}
//
//func createTemplateCacheOld(t string) error {
//	templates := []string{
//		fmt.Sprintf("./templates/%s", t),
//		"./templates/base.layout.tmpl",
//	}
//
//	// parse the template
//	tmpl, err := template.ParseFiles(templates...)
//	if err != nil {
//		return err
//	}
//
//	// add template to cache
//	tc[t] = tmpl
//
//	return nil
//}

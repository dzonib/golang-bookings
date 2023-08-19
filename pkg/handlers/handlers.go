package handlers

import (
	"net/http"

	"github.com/dzonib/golang-app-with-templates/pkg/models"

	"github.com/dzonib/golang-app-with-templates/pkg/config"

	"github.com/dzonib/golang-app-with-templates/pkg/templates"
)

// Repo is the repository used by handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr

	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)

	templates.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	// preform some logic
	stringMap := map[string]string{
		"test":     "Hello bruh",
		"remoteIP": remoteIP,
	}

	// send data to template
	templates.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

package handlers

import (
	"github.com/EraldBa/webApp/pkg/models"
	"net/http"

	"github.com/EraldBa/webApp/pkg/config"
	"github.com/EraldBa/webApp/pkg/render"
)

// Repository is the prototype of repository type
type Repository struct {
	App *config.AppConfig
}

// Repo is the handles repository
var Repo *Repository

// NewRepo makes new repo
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repo for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) HomeHandler(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.gohtml", &models.TemplateData{})
}

// AboutHandler handles /about requests
func (m *Repository) AboutHandler(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["test"] = "Hello from backend!"
	stringMap["remote_ip"] = remoteIP
	render.RenderTemplate(w, "about.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}

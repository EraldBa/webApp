package handlers

import (
	"github.com/EraldBa/webApp/pkg/models"
	"net/http"

	"github.com/EraldBa/webApp/pkg/config"
	"github.com/EraldBa/webApp/pkg/render"

	_ "github.com/go-sql-driver/mysql"
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
	render.RenderTemplate(w, r, "home.page.gohtml", &models.TemplateData{})
}

// AboutHandler handles /about requests
func (m *Repository) AboutHandler(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["test"] = "Hello from backend!"
	stringMap["remote_ip"] = remoteIP
	render.RenderTemplate(w, r, "about.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}

//func OpenDB() *sql.DB {
//	db, err := sql.Open("mysql", "root:nyc595486672@tcp(127.0.0.1:3306)/dates")
//
//	if err != nil {
//		log.Fatal(err.Error())
//	}
//	return db
//}

func (m *Repository) UpdateCalHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "update-cal.page.gohtml", &models.TemplateData{})
}

func (m *Repository) PostUpdateCalHandler(w http.ResponseWriter, r *http.Request) {
	calStats := map[string]string{
		"breakfast": r.Form.Get("breakfast"),
		"lunch":     r.Form.Get("lunch"),
		"dinner":    r.Form.Get("dinner"),
		"snacks":    r.Form.Get("snacks"),
	}

	render.RenderTemplate(w, r, "show-stats.page.gohtml", &models.TemplateData{
		StringMap: calStats,
	})
}

func (m *Repository) AboutPost(w http.ResponseWriter, r *http.Request) {

	userMap := make(map[string]string)
	userMap["user"] = r.Form.Get("username")
	userMap["password"] = r.Form.Get("password")

	render.RenderTemplate(w, r, "success-signup.page.gohtml", &models.TemplateData{
		StringMap: userMap,
	})

}

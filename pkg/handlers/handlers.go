package handlers

import (
	"encoding/json"
	"github.com/EraldBa/webApp/pkg/models"
	"github.com/justinas/nosurf"
	"log"
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

func (m *Repository) DashboardHandler(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{
		"breakfast": "100",
		"lunch":     "100",
		"dinner":    "100",
		"snacks":    "100",
		"protein":   "200",
		"carbs":     "400",
		"fats":      "100",
	}
	render.RenderTemplate(w, r, "dashboard.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) PostDashboardHandler(w http.ResponseWriter, r *http.Request) {

	stringMap := map[string]string{
		"time_of_day":  r.Form.Get("time_of_day"),
		"desired_date": r.Form.Get("desired_date"),
		"breakfast":    r.Form.Get("calorie"),
		"lunch":        r.Form.Get("protein"),
		"dinner":       r.Form.Get("carbs"),
		"snacks":       r.Form.Get("fats"),
	}
	log.Println(stringMap)
}

type dateRecievedJSON struct {
	Date      string `json:"date"`
	CSRFToken string `json:"csrf_token"`
}
type dateResponseJSON struct {
	Breakfast int `json:"breakfast"`
	Lunch     int `json:"lunch"`
	Dinner    int `json:"dinner"`
	Snacks    int `json:"snacks"`
	Protein   int `json:"protein"`
	Carbs     int `json:"carbs"`
	Fats      int `json:"fats"`
}

func (m *Repository) PostDashNewHandler(w http.ResponseWriter, r *http.Request) {
	var p dateRecievedJSON

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		log.Fatal(err)
	}
	if !nosurf.VerifyToken(nosurf.Token(r), p.CSRFToken) {
		_, _ = w.Write([]byte("Error 400. Server refused Connection"))
		return
	}

	if p.Date == "2022-11-19" {
		a := dateResponseJSON{
			Breakfast: 400,
			Lunch:     500,
			Dinner:    600,
			Snacks:    200,
			Protein:   180,
			Carbs:     360,
			Fats:      80,
		}
		b, _ := json.Marshal(a)
		_, _ = w.Write(b)
	}
}

func (m *Repository) MemberHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "member.page.gohtml", &models.TemplateData{
		Error: "Failed Log In Attempt!",
	})
}

func (m *Repository) PostSignUpHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)

}

func (m *Repository) PostLogInHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func (m *Repository) ContactHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.gohtml", &models.TemplateData{})
}

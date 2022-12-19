package handlers

import (
	"encoding/json"
	"github.com/EraldBa/webApp/pkg/driver"
	"github.com/EraldBa/webApp/pkg/models"
	"github.com/EraldBa/webApp/pkg/repository"
	"github.com/EraldBa/webApp/pkg/repository/dbrepo"
	"github.com/justinas/nosurf"
	"log"
	"net/http"

	"github.com/EraldBa/webApp/pkg/config"
	"github.com/EraldBa/webApp/pkg/render"
)

// Repository is the prototype of repository type
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// Repo is the handle's repository
var Repo *Repository

// NewRepo makes new repo
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}
}

// NewHandlers sets the repo for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) HomeHandler(w http.ResponseWriter, r *http.Request) {
	//remoteIP := r.RemoteAddr
	//m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.Template(w, r, "home.page.gohtml", &models.TemplateData{})
}

// AboutHandler handles /about requests
func (m *Repository) AboutHandler(w http.ResponseWriter, r *http.Request) {
	//stringMap := make(map[string]string)
	//remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	//
	//stringMap["test"] = "Hello from backend!"
	//stringMap["remote_ip"] = remoteIP
	render.Template(w, r, "about.page.gohtml", &models.TemplateData{})
}

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
	render.Template(w, r, "dashboard.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) PostDashboardHandler(w http.ResponseWriter, r *http.Request) {
	stats := models.StatsGet{
		TimeOfDay: r.Form.Get("time_of_day"),
		Date:      r.Form.Get("desired_date"),
		Calories:  r.Form.Get("calorie"),
		Protein:   r.Form.Get("protein"),
		Carbs:     r.Form.Get("carbs"),
		Fats:      r.Form.Get("fats"),
		UserID:    m.App.Session.GetString(r.Context(), "user_id"),
	}

	if err := m.DB.CheckStats(stats.Date); err != nil {
		if err = m.DB.UpdateStats(&stats); err != nil {
			log.Println("Die")
		}
	} else {
		if err = m.DB.InsertNewStats(&stats); err != nil {
			log.Println("Die")
		}
	}

	log.Println(stats)
}

func (m *Repository) PostDashRefreshHandler(w http.ResponseWriter, r *http.Request) {
	var p models.GetDate

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		log.Println(err)
		return
	}
	if !nosurf.VerifyToken(nosurf.Token(r), p.CSRFToken) {
		_, _ = w.Write([]byte("Error 400. Server refused Connection"))
		return
	}

	if p.Date == "2022-11-19" {
		a := models.StatsSend{
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
	render.Template(w, r, "member.page.gohtml", &models.TemplateData{})
}

func (m *Repository) PostSignUpHandler(w http.ResponseWriter, r *http.Request) {
	signupData := models.User{
		Username: r.Form.Get("username"),
		Email:    r.Form.Get("email"),
		Password: r.Form.Get("password"),
	}
	log.Println(signupData)
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)

}

func (m *Repository) PostLogInHandler(w http.ResponseWriter, r *http.Request) {

	loginData := models.User{
		Username: r.Form.Get("username"),
		Password: r.Form.Get("password"),
	}

	if loginData.Username == "Erald" {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	render.Template(w, r, "member.page.gohtml", &models.TemplateData{
		Error: "Login unsuccessful, check your info and try again",
	})

}

func (m *Repository) ContactHandler(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "contact.page.gohtml", &models.TemplateData{})
}

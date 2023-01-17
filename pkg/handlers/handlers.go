package handlers

import (
	"encoding/json"
	"github.com/EraldBa/webApp/pkg/config"
	"github.com/EraldBa/webApp/pkg/driver"
	"github.com/EraldBa/webApp/pkg/helpers"
	"github.com/EraldBa/webApp/pkg/models"
	"github.com/EraldBa/webApp/pkg/render"
	"github.com/EraldBa/webApp/pkg/repository"
	"github.com/EraldBa/webApp/pkg/repository/dbrepo"
	"github.com/justinas/nosurf"
	"net/http"
	"time"
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
	render.Template(w, r, "home.page.gohtml", &models.TemplateData{})
}

// AboutHandler handles /about requests
func (m *Repository) AboutHandler(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "about.page.gohtml", &models.TemplateData{})
}

func (m *Repository) DashboardHandler(w http.ResponseWriter, r *http.Request) {
	userID := m.App.Session.GetInt(r.Context(), "user_id")
	isLoggedIn := helpers.IsAuthenticated(r)

	if !isLoggedIn {
		m.App.Session.Put(r.Context(), "error", "You're not logged in, please log in to view your stats!")
		http.Redirect(w, r, "/member", http.StatusSeeOther)
		return
	}

	currentDate := time.Now().Format("2006-01-02")
	statsSend := m.DB.GetStats(currentDate, userID)

	floatMap := map[string]float32{
		"breakfast": statsSend.Breakfast,
		"lunch":     statsSend.Lunch,
		"dinner":    statsSend.Dinner,
		"snacks":    statsSend.Snacks,
		"protein":   statsSend.Protein,
		"carbs":     statsSend.Carbs,
		"fats":      statsSend.Fats,
	}

	render.Template(w, r, "dashboard.page.gohtml", &models.TemplateData{
		FloatMap: floatMap,
	})
}

func (m *Repository) PostDashboardHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	userID := m.App.Session.GetInt(r.Context(), "user_id")

	macros := models.Macros{
		Request:   r,
		Precision: 2,
		BitSize:   32,
	}

	stats := models.StatsGet{
		TimeOfDay: r.Form.Get("time_of_day"),
		Date:      r.Form.Get("desired_date"),
		Calories:  macros.GetMacro("calorie"),
		Protein:   macros.GetMacro("protein"),
		Carbs:     macros.GetMacro("carbs"),
		Fats:      macros.GetMacro("fats"),
		UserID:    userID,
	}
	// If there's an error, row doesn't exist so making a new one, else update the row
	if err = m.DB.CheckStats(stats.Date, userID); err == nil {
		err = m.DB.UpdateStats(&stats)
	} else {
		err = m.DB.InsertNewStats(&stats)
	}

	if err != nil {
		helpers.ServerError(w, err)
	}
}

func (m *Repository) PostDashRefreshHandler(w http.ResponseWriter, r *http.Request) {
	var receivedJSON models.GetDate
	// Receiving json from frontend
	if err := json.NewDecoder(r.Body).Decode(&receivedJSON); err != nil {
		helpers.ServerError(w, err)
		return
	}
	// Checking if token sent in json is valid
	if !nosurf.VerifyToken(nosurf.Token(r), receivedJSON.CSRFToken) {
		_, _ = w.Write([]byte("Error 400. Server refused Connection"))
		return
	}
	userID := m.App.Session.GetInt(r.Context(), "user_id")

	statsSend := m.DB.GetStats(receivedJSON.Date, userID)

	statsSendJSON, err := json.Marshal(statsSend)
	if err != nil {
		helpers.ServerError(w, err)
	}

	if _, err = w.Write(statsSendJSON); err != nil {
		helpers.ServerError(w, err)
	}
}

func (m *Repository) MemberHandler(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "member.page.gohtml", &models.TemplateData{})
}

func (m *Repository) PostSignUpHandler(w http.ResponseWriter, r *http.Request) {
	signupData := models.User{
		Username:    r.Form.Get("username"),
		Email:       r.Form.Get("email"),
		Password:    r.Form.Get("password"),
		AccessLevel: 1,
	}

	if err := m.DB.InsertUser(&signupData); err != nil {
		helpers.ClientError(w, http.StatusInternalServerError)
		return
	}

	m.PostLogInHandler(w, r)
}

func (m *Repository) PostLogInHandler(w http.ResponseWriter, r *http.Request) {
	_ = m.App.Session.RenewToken(r.Context())

	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	username := r.Form.Get("username")
	password := r.Form.Get("password")
	id, _, err := m.DB.Authenticator(username, password)
	if err != nil {
		m.App.Session.Put(r.Context(), "error", err.Error())
		http.Redirect(w, r, "/member", http.StatusSeeOther)
		return
	}
	m.App.Session.Put(r.Context(), "flash", "Logged in successfully!")
	m.App.Session.Put(r.Context(), "user_id", id)
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func (m *Repository) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	err := m.App.Session.Destroy(r.Context())
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	err = m.App.Session.RenewToken(r.Context())
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	m.App.Session.Put(r.Context(), "flash", "Logged out successfully!")

	http.Redirect(w, r, "/member", http.StatusSeeOther)
}

func (m *Repository) ContactHandler(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "contact.page.gohtml", &models.TemplateData{})
}

package handlers

import (
	"encoding/json"
	"fmt"
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
	isLoggedIn := helpers.IsAuthenticated(r)
	if !isLoggedIn {
		m.App.Session.Put(r.Context(), "error", "You're not logged in, please log in to view your stats!")
		http.Redirect(w, r, "/member", http.StatusSeeOther)
		return
	}

	userID := m.App.Session.Get(r.Context(), "user_id").(uint)
	currentDate := time.Now().Format("2006-01-02")
	statsSend := m.DB.GetStats(currentDate, userID)

	render.Template(w, r, "dashboard.page.gohtml", &models.TemplateData{
		Stats: statsSend,
	})
}

func (m *Repository) PostDashboardHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	userID := m.App.Session.Get(r.Context(), "user_id").(uint)

	stats := models.StatsGet{
		TimeOfDay: r.Form.Get("time_of_day"),
		Date:      r.Form.Get("desired_date"),
		UserID:    userID,
	}
	stats.SetMacros(r)

	noRows := m.DB.CheckStats(stats.Date, userID)
	if noRows {
		err = m.DB.InsertNewStats(&stats)
	} else {
		err = m.DB.UpdateStats(&stats)
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

	csrfIsValid := nosurf.VerifyToken(nosurf.Token(r), receivedJSON.CSRFToken)
	if !csrfIsValid {
		helpers.ClientError(w, http.StatusBadRequest)
		return
	}

	userID := m.App.Session.Get(r.Context(), "user_id").(uint)
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
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	signupData := models.User{
		Username: r.Form.Get("username"),
		Email:    r.Form.Get("email"),
		Password: r.Form.Get("password"),
	}

	if err := m.DB.InsertUser(&signupData); err != nil {
		m.App.Session.Put(r.Context(), "error", "Username and/or email are already taken!")
		http.Redirect(w, r, "/member", http.StatusSeeOther)
		return
	}
	m.App.Session.Put(r.Context(), "new_user", true)

	http.Redirect(w, r, "/logged-in", http.StatusTemporaryRedirect)
}

func (m *Repository) PostLogInHandler(w http.ResponseWriter, r *http.Request) {
	var msgKey string
	var msg string

	_ = m.App.Session.RenewToken(r.Context())

	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	id, err := m.DB.Authenticator(username, password)
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "Wrong credentials. Try again.")
		http.Redirect(w, r, "/member", http.StatusSeeOther)
		return
	}

	newUser := m.App.Session.PopBool(r.Context(), "new_user")
	if newUser {
		msgKey = "success"
		msg = fmt.Sprintf("Welcome to FitBuddy %s!", username)
	} else {
		msgKey = "flash"
		msg = fmt.Sprintf("Logged in successfully as %s!", username)
	}

	m.App.Session.Put(r.Context(), msgKey, msg)
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

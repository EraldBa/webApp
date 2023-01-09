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
	"log"
	"net/http"
	"strconv"
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
	//remoteIP := r.RemoteAddr
	//m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	m.App.Session.Put(r.Context(), "user_id", 1)
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
	userID := m.App.Session.GetInt(r.Context(), "user_id")
	log.Println(userID)
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
	log.Println(floatMap)
	render.Template(w, r, "dashboard.page.gohtml", &models.TemplateData{
		FloatMap: floatMap,
	})
}

func (m *Repository) PostDashboardHandler(w http.ResponseWriter, r *http.Request) {
	bitSize := 32
	userID := m.App.Session.GetInt(r.Context(), "user_id")
	calories, err := strconv.ParseFloat(r.Form.Get("calorie"), bitSize)
	helpers.ErrorCheck(err)
	protein, err := strconv.ParseFloat(r.Form.Get("protein"), bitSize)
	helpers.ErrorCheck(err)
	carbs, err := strconv.ParseFloat(r.Form.Get("carbs"), bitSize)
	helpers.ErrorCheck(err)
	fats, err := strconv.ParseFloat(r.Form.Get("fats"), bitSize)
	helpers.ErrorCheck(err)

	stats := models.StatsGet{
		TimeOfDay: r.Form.Get("time_of_day"),
		Date:      r.Form.Get("desired_date"),
		Calories:  calories,
		Protein:   protein,
		Carbs:     carbs,
		Fats:      fats,
		UserID:    userID,
	}
	// If there's an error, row doesn't exist so making a new one, else update the row
	if err = m.DB.CheckStats(stats.Date, userID); err != nil {
		err = m.DB.UpdateStats(&stats)
		helpers.ErrorCheck(err)
	} else {
		err = m.DB.InsertNewStats(&stats)
		helpers.ErrorCheck(err)
	}
}

func (m *Repository) PostDashRefreshHandler(w http.ResponseWriter, r *http.Request) {
	var receivedJSON models.GetDate
	// Receiving json from frontend
	if err := json.NewDecoder(r.Body).Decode(&receivedJSON); err != nil {
		log.Println(err)
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
	helpers.ErrorCheck(err)

	_, _ = w.Write(statsSendJSON)

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

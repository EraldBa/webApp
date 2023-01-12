package models

import "net/http"

// User holds the users info
type User struct {
	ID          int
	Username    string
	Email       string
	Password    string
	AccessLevel int
}

// StatsGet is the stats info that the backend needs
// to insert or update stats in database
type StatsGet struct {
	Date      string
	TimeOfDay string
	Calories  float64
	Protein   float64
	Carbs     float64
	Fats      float64
	UserID    int
}

// StatsSend holds the info that needs to be sent to frontend user
type StatsSend struct {
	Breakfast float32 `json:"breakfast"`
	Lunch     float32 `json:"lunch"`
	Dinner    float32 `json:"dinner"`
	Snacks    float32 `json:"snacks"`
	Protein   float32 `json:"protein"`
	Carbs     float32 `json:"carbs"`
	Fats      float32 `json:"fats"`
}

// GetDate is the json that backend receives when frontend user requests stat data
type GetDate struct {
	Date      string `json:"date"`
	CSRFToken string `json:"csrf_token"`
}

type Macros struct {
	Request   *http.Request
	Precision uint8
	BitSize   int
}

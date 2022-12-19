package models

import "time"

type User struct {
	ID          int
	Username    string
	Email       string
	Password    string
	AccessLevel int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Login struct {
	Username string
	Password string
}

type StatsForm struct {
	Date      string
	TimeOfDay string
	Calories  string
	Protein   string
	Carbs     string
	Fats      string
}

type StatsSend struct {
	Breakfast int `json:"breakfast"`
	Lunch     int `json:"lunch"`
	Dinner    int `json:"dinner"`
	Snacks    int `json:"snacks"`
	Protein   int `json:"protein"`
	Carbs     int `json:"carbs"`
	Fats      int `json:"fats"`
}

type GetDate struct {
	Date      string `json:"date"`
	CSRFToken string `json:"csrf_token"`
}

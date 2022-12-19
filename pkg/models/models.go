package models

type User struct {
	ID          int
	Username    string
	Email       string
	Password    string
	AccessLevel int
}

type StatsGet struct {
	Date      string
	TimeOfDay string
	Calories  string
	Protein   string
	Carbs     string
	Fats      string
	UserID    string
}

type StatsSend struct {
	Breakfast float32 `json:"breakfast"`
	Lunch     float32 `json:"lunch"`
	Dinner    float32 `json:"dinner"`
	Snacks    float32 `json:"snacks"`
	Protein   float32 `json:"protein"`
	Carbs     float32 `json:"carbs"`
	Fats      float32 `json:"fats"`
}

type GetDate struct {
	Date      string `json:"date"`
	CSRFToken string `json:"csrf_token"`
}

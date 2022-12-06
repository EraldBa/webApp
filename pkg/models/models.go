package modelsa

type Signup struct {
	Username string
	Email    string
	Password string
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

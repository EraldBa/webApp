package models

import (
	"log"
	"net/http"
	"strconv"
)

var macros = [4]string{
	"calorie",
	"protein",
	"carbs",
	"fats",
}

// SetMacros takes strings from form data, converts them to floats and sets the macros of StatsGet
func (s *StatsGet) SetMacros(r *http.Request) {
	values := make(map[string]float32, 4)

	for _, macro := range macros {
		value, err := strconv.ParseFloat(r.Form.Get(macro), 32)
		if err != nil {
			log.Println(err)
			continue
		}
		values[macro] = float32(value)
	}

	s.Calories = values["calorie"]
	s.Protein = values["protein"]
	s.Carbs = values["carbs"]
	s.Fats = values["fats"]
}

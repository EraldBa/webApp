package models

import (
	"log"
	"net/http"
	"strconv"
)

var macros = map[string]float32{
	"calorie": 0,
	"protein": 0,
	"carbs":   0,
	"fats":    0,
}

// SetMacros takes strings from form data, converts them to floats and sets the macros of StatsGet
func (s *StatsGet) SetMacros(r *http.Request) {
	for macro, _ := range macros {
		value, err := strconv.ParseFloat(r.Form.Get(macro), 32)
		if err != nil {
			macros[macro] = 0
			log.Println(err)
			continue
		}
		macros[macro] = float32(value)
	}

	s.Calories = macros["calorie"]
	s.Protein = macros["protein"]
	s.Carbs = macros["carbs"]
	s.Fats = macros["fats"]
}

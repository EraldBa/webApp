package models

import (
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
	for _, formField := range macros {
		value, _ := strconv.ParseFloat(r.Form.Get(formField), 32)
		value32 := float32(value)

		switch formField {
		case "calorie":
			s.Calories = value32
		case "protein":
			s.Protein = value32
		case "carbs":
			s.Carbs = value32
		case "fats":
			s.Fats = value32
		}
	}
}

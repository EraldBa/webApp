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
	for i, formField := range macros {
		value, _ := strconv.ParseFloat(r.Form.Get(formField), 32)
		value32 := float32(value)

		switch i {
		case 0:
			s.Calories = value32
		case 1:
			s.Protein = value32
		case 2:
			s.Carbs = value32
		case 3:
			s.Fats = value32
		}
	}
}

package models

import (
	"fmt"
	"net/http"
)

var macros = [4]string{
	"calorie",
	"protein",
	"carbs",
	"fats",
}

// SetMacros takes strings from form data, converts them to floats and sets the macros of StatsGet
func (s *StatsGet) SetMacros(r *http.Request) {
	var value float32
	var err error

	for _, formField := range macros {
		_, err = fmt.Sscanf(r.Form.Get(formField), "%f", &value)
		if err != nil {
			value = 0
		}

		switch formField {
		case "calorie":
			s.Calories = value
		case "protein":
			s.Protein = value
		case "carbs":
			s.Carbs = value
		case "fats":
			s.Fats = value
		}
	}
}

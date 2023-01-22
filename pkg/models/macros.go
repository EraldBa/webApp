package models

import (
	"net/http"
	"reflect"
	"strconv"
)

var macros = map[string]string{
	"calorie": "Calories",
	"protein": "Protein",
	"carbs":   "Carbs",
	"fats":    "Fats",
}

// SetMacros takes strings from form data, converts them to floats and sets the macros of StatsGet
func (s *StatsGet) SetMacros(r *http.Request) {
	statsVal := reflect.ValueOf(s).Elem()
	for formField, structField := range macros {
		value, _ := strconv.ParseFloat(r.Form.Get(formField), 32)
		statsVal.FieldByName(structField).SetFloat(value)
	}
}

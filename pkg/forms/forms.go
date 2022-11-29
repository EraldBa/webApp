package forms

import (
	"net/http"
	"net/url"
)

// Form is a custom from struct, has url.Values object
type Form struct {
	url.Values
	Errors errors
}

// New initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		map[string][]string{},
	}
}

// Has checks if form field has values
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		return false
	}
	return true
}

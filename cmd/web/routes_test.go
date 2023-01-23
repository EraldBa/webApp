package main

import (
	"testing"

	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {
	mux := routes()

	switch mux.(type) {
	case *chi.Mux:
	// All is good
	default:
		t.Error("Type is not *chi.Mux")
	}
}

package main

import (
	"github.com/EraldBa/webApp/pkg/config"
	"github.com/go-chi/chi"
	"testing"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig
	mux := routes(&app)

	switch mux.(type) {
	case *chi.Mux:
	// All is good
	default:
		t.Error("Type is not *chi.Mux")
	}
}

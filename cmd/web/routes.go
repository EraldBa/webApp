package main

import (
	"github.com/EraldBa/webApp/pkg/config"
	"github.com/EraldBa/webApp/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.HomeHandler)
	mux.Get("/about", handlers.Repo.AboutHandler)
	mux.Post("/about", handlers.Repo.AboutPost)
	mux.Get("/update-calories", handlers.Repo.UpdateCalHandler)
	mux.Post("/show-stats", handlers.Repo.PostUpdateCalHandler)

	return mux
}

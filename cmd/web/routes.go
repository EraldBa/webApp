package main

import (
	"github.com/EraldBa/webApp/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.HomeHandler)
	mux.Get("/about", handlers.Repo.AboutHandler)
	mux.Get("/dashboard", handlers.Repo.DashboardHandler)
	mux.Post("/dashboard", handlers.Repo.PostDashboardHandler)
	mux.Post("/dashboard-new", handlers.Repo.PostDashRefreshHandler)
	mux.Get("/member", handlers.Repo.MemberHandler)
	mux.Post("/signed-up", handlers.Repo.PostSignUpHandler)
	mux.Post("/logged-in", handlers.Repo.PostLogInHandler)
	mux.Get("/logout", handlers.Repo.LogoutHandler)
	mux.Get("/contact", handlers.Repo.ContactHandler)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}

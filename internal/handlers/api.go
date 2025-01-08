package handlers

import (
	"github.com/Drake-Programming/go-currency-api/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(route *chi.Mux) {
	// Global middleware
	route.Use(chimiddle.StripSlashes)

	route.Route("/account", func(router chi.Router) {

		// Middleware for /account route
		router.Use(middleware.Authorization)

		router.Get("/balance", GetBalance)
	})
}

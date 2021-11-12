package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
)

func (s *Server) middlewares() error {

	addCorsMiddleware(s.router)
	return nil
}

func addCorsMiddleware(router *chi.Mux) {
	allowedOrigins := []string{"*"}

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	router.Use(corsMiddleware.Handler)
}

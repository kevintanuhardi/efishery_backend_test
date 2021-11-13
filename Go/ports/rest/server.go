package rest

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kevintanuhardi/efishery_backend_test/config"
	"gorm.io/gorm"
)

type Server struct {
	cfg    *config.Config
	router *chi.Mux
	db     *gorm.DB
}

func (s *Server) Run() error {
	// Use Middlewares
	err := s.middlewares()
	if err != nil {
		return err
	}

	// Use Routes
	s.routes()

	return http.ListenAndServe(fmt.Sprintf(":%d", s.cfg.Port.HTTP), s.router)
}

package rest

import (
	"net/http"

	"github.com/kevintanuhardi/efishery_backend_test/pkg/render"
)

func (s *Server) routes() {
	r := s.router

	// ctx := context.Background()

	r.Get("/ping", func (w http.ResponseWriter, r *http.Request) {
		render.Response(w, http.StatusOK, "orders", render.EmptyResponse, render.EmptyResponse)
	})

}

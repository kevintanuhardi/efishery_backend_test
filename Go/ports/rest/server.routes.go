package rest

import (
	"context"
	"net/http"

	"github.com/kevintanuhardi/efishery_backend_test/domain/user/repository/mysql"
	"github.com/kevintanuhardi/efishery_backend_test/domain/user/usecase"
	"github.com/kevintanuhardi/efishery_backend_test/pkg/render"
)

func (s *Server) routes() {
	r := s.router

	ctx := context.Background()

	r.Get("/ping", func (w http.ResponseWriter, r *http.Request) {
		render.Response(w, http.StatusOK, "pong", render.EmptyResponse, render.EmptyResponse)
	})

	user := &User{
		Usecase: usecase.NewService(
			mysql.NewRepository(s.db),
		),
	}
	user.Register(ctx, r)	

}

package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kevintanuhardi/efishery_backend_test/domain/user/entity"
	"github.com/kevintanuhardi/efishery_backend_test/domain/user/usecase"
	"github.com/kevintanuhardi/efishery_backend_test/pkg/render"
	"github.com/kevintanuhardi/efishery_backend_test/pkg/stringHelper"
)

type User struct {
	Usecase *usecase.Service
}

func (u *User) Register(ctx context.Context, r chi.Router) {
	r.Route("/user", func(r chi.Router) {
		r.Post("/register", u.registerUser)
	})
}

type registerUserResponse struct {
	Password	string			`json:"password"`
}

func (u *User) registerUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	decoder := json.NewDecoder(r.Body)
	var request usecase.CreateUserRequest
	if err := decoder.Decode(&request); err != nil {
		render.Response(w, http.StatusNotAcceptable, render.EmptyResponse, render.EmptyResponse, render.EmptyResponse)
		return
	}

	password := stringHelper.RandStringRunes(4)

	fmt.Println(password)

	_, err := u.Usecase.CreateUser(ctx, entity.User{
		Name: request.Name,
		Phone: request.Phone,
		Role: request.Role,
		Password: password,
	});
	if err != nil {
		render.Response(w, http.StatusNotAcceptable, render.EmptyResponse, render.EmptyResponse, render.EmptyResponse)
		return
	}

	response := registerUserResponse{
		Password: password,
	}

	render.Response(w, http.StatusOK, response, render.EmptyResponse, render.EmptyResponse)
}

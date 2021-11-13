package rest

import (
	"context"
	"encoding/json"
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
		r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
			u.login(w , r.WithContext(ctx))
		})
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


	_, err := u.Usecase.CreateUser(ctx, entity.User{
		Name: request.Name,
		Phone: request.Phone,
		Role: request.Role,
		Password: password,
	});
	if err != nil {
		if err.Error() == "duplicate_user" {
			render.Response(w, http.StatusBadRequest, render.EmptyResponse, "This phone number is already registered", render.EmptyResponse)
		} else {
			render.Response(w, http.StatusNotAcceptable, render.EmptyResponse, render.EmptyResponse, render.EmptyResponse)
		}
		return
	}

	response := registerUserResponse{
		Password: password,
	}

	render.Response(w, http.StatusOK, response, render.EmptyResponse, render.EmptyResponse)
}

func (u *User) login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()


	decoder := json.NewDecoder(r.Body)
	var request usecase.LoginUserRequest
	if err := decoder.Decode(&request); err != nil {
		render.Response(w, http.StatusNotAcceptable, render.EmptyResponse, render.EmptyResponse, render.EmptyResponse)
		return
	}


	_, err := u.Usecase.Login(ctx, request);
	if err != nil {
		render.Response(w, http.StatusUnauthorized, render.EmptyResponse, err.Error(), render.EmptyResponse)
		return
	}

	render.Response(w, http.StatusOK, render.EmptyResponse, render.EmptyResponse, render.EmptyResponse)
}

package rest

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt"
	"github.com/kevintanuhardi/efishery_backend_test/domain/user/entity"
	"github.com/kevintanuhardi/efishery_backend_test/domain/user/usecase"
	"github.com/kevintanuhardi/efishery_backend_test/pkg/render"
	"github.com/kevintanuhardi/efishery_backend_test/pkg/stringHelper"
)

type User struct {
	Usecase *usecase.Service
}

func (u *User) Register(ctx context.Context, r chi.Router) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", u.registerUser)
		r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
			u.login(w , r.WithContext(ctx))
		})
		r.With(ExternalAccess(false)).Get("/token/introspect", u.tokenIntrospect)
	})
}

type registerUserResponse struct {
	Password	string			`json:"password"`
}

type tokenIntrospectResponse struct {
	Name string `json:"name"`
	Phone string `json:"phone"`
	Role string `json:"role"`
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


	token, err := u.Usecase.Login(ctx, request);
	if err != nil {
		render.Response(w, http.StatusUnauthorized, render.EmptyResponse, err.Error(), render.EmptyResponse)
		return
	}

	render.Response(w, http.StatusOK, token, render.EmptyResponse, render.EmptyResponse)
}

func (u *User) tokenIntrospect(w http.ResponseWriter, r *http.Request) {

	privateClaim := r.Context().Value("privateClaim").(jwt.MapClaims)

	claim := tokenIntrospectResponse{
		Name: privateClaim["name"].(string),
		Phone: privateClaim["phone"].(string),
		Role: privateClaim["role"].(string),
	}
	render.Response(w, http.StatusOK, claim , render.EmptyResponse, render.EmptyResponse)
}
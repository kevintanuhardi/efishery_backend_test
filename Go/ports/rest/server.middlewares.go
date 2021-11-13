package rest

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt"
	"github.com/kevintanuhardi/efishery_backend_test/config"
	"github.com/rs/cors"
)

var JWT_SIGNING_METHOD = jwt.SigningMethodHS256

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

func ExternalAccess(adminOnly bool) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			authString := r.Header.Get("Authorization")
			if authString == "" {
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				return
			}
			authArr := strings.Split(authString, " ")
			if len(authArr) != 2 { //strings should consist of 2 words
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				return
			}
			if authArr[0] != "Bearer" { //auth type should be Bearer
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				return
			}
			if authArr[1] == "" { // token should not be empty
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				return
			}

			cfg, _ := config.Load("config/config.json")
			jwtSecret := cfg.JWT_SECRET
			newToken, _ := jwt.Parse(authArr[1], func(token *jwt.Token) (interface{}, error) {
				if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Signing method invalid")
				} else if method != JWT_SIGNING_METHOD {
						return nil, fmt.Errorf("Signing method invalid")
				}
			
				return []byte(jwtSecret), nil
			})
			claims, ok := newToken.Claims.(jwt.MapClaims)
			if !ok {
			// if !ok || !newToken.Valid {
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
					return
			}

			if adminOnly && claims["role"] != "admin" {
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			}

			ctx1 := context.WithValue(ctx, "privateClaim", claims)

			next.ServeHTTP(w, r.WithContext(ctx1))
		})
	}
}
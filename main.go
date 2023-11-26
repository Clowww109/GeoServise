package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"main/authentication"
	_ "main/docs"
	"main/services"
	"net/http"
)

// @title GeoService
// @version 1.0
// @description Api server for GeoServise

// @host localhost:8080
// @BasePath /
func main() {

	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Post("/api/register/", authentication.UserRegister)
		r.Post("/api/login/", authentication.UserLogin)
	})

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(authentication.TokenAuth))

		r.Use(jwtauth.Authenticator(authentication.TokenAuth))

		r.Post("/api/address/search/", services.GetAddress)

		r.Post("/api/address/geocode/", services.GetGeocode)
	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	http.ListenAndServe(":8080", r)
}

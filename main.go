package main

import (
	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
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

	r.Post("/api/address/search", services.GetAddress)

	r.Post("/api/address/geocode", services.GetGeocode)

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
	))

	http.ListenAndServe(":8080", r)
}

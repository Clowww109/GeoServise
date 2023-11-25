package main

import (
	"github.com/go-chi/chi"
	"main/GeoServise/service"
	"net/http"
)

// @title GeoServise
// @version 1.0
// @description Api server for GeoServise

// @host localhost:8080
// @BasePath /

func main() {

	r := chi.NewRouter()

	r.Post("/api/address/search", service.GetAddress)

	r.Post("/api/address/geocode", service.GetGeocode)

	r.Get("/swagger/*", service.HandleSwagger)

	http.ListenAndServe(":8080", r)
}

package main

import (
	_ "github.com/SergeiKaluckij/minitwitter-backend/docs"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"net/http"
)

// Swagger Settings
//
//	@title			Minitwitter API
//	@version		1.0.0
//	@description	Description
//	@host			localhost:9001
//	@BasePath		/api/v1
func main() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.CleanPath)
	router.Use(middleware.Recoverer)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:9001/swagger/doc.json"), //The url pointing to API definition
	))

	http.ListenAndServe(":9001", router)
}

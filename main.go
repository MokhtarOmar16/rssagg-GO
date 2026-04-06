package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(os.Getenv("PORT"))
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"*"},
		MaxAge:         300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", HandlerReadiness)
	v1Router.Get("/err", HandlerErr)
	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + os.Getenv("PORT"),
	}

	log.Print("Starting server on port " + os.Getenv("PORT"))
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server")
	}
}

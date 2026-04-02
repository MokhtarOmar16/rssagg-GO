package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	// "github.com/go-chi/cors"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(os.Getenv("PORT"))
	router := chi.NewRouter()
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

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MokhtarOmar16/rssagg-GO/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)


type apiconfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	fmt.Println(os.Getenv("PORT"))
	dbURL := os.Getenv("DB_URL")
	fmt.Println(dbURL)
	conn, err := sql.Open("postgres", dbURL)
	if err != nil{
		log.Fatal(err)
	}
	
	apiCfg :=apiconfig{
		DB :database.New(conn),
	}



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
	router.Post("/users", apiCfg.CreateUserHandler)
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

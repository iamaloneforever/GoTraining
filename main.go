package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/iamaloneforever/GoTraining/db"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *db.Queries
}

func main() {
	godotenv.Load(".env")

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL not found")
	}

	portString := os.Getenv("PORT")
	if portString == "" {
		portString = "8080"
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("DB cannot connect: %v", err)
	}

	apiCon := apiConfig{
		DB: db.New(conn),
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*", "https://*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/ready", handleReadiness)
	v1Router.Post("/users", apiCon.handleCreateUser)

	router.Mount("/v1", v1Router)

	log.Printf("Server starting on port %s", portString)
	log.Fatal(http.ListenAndServe(":"+portString, router))
}

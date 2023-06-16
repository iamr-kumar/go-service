package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/iamr-kumar/go-service/internal/databases"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *databases.Queries
}

func main() {
	
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Could not connect to the database")
	}

	apiCfg := apiConfig {
		DB: databases.New(conn),
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"*"},
			AllowCredentials: false,
			ExposedHeaders:   []string{"Link"},
			MaxAge:           300,
		}))

		v1Router := chi.NewRouter()

		v1Router.Get("/healthz", handlerReadiness)
		v1Router.Get("/error", handleError)
		v1Router.Post(("/users"), apiCfg.handleCreateUser)

		router.Mount("/api/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}


	log.Println("Server is running on port " + portString)

	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
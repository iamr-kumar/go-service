package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
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

		router.Mount("/api/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}


	log.Println("Server is running on port " + portString)

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
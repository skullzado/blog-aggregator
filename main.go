package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{}))

	subRouter := chi.NewRouter()
	subRouter.Get("/readiness", handlerReadiness)
	subRouter.Get("/err", errorHandler)

	router.Mount("/v1", subRouter)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Printf("Serving on port: %s", port)
	log.Fatal(srv.ListenAndServe())
}

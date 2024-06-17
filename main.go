package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT must be set")
	}

	router := chi.NewRouter()
	v1Router := chi.NewRouter()
	router.Mount("/v1", v1Router)

	v1Router.HandleFunc("GET /healthz", handlerReadiness)
	v1Router.HandleFunc("GET /err", func(w http.ResponseWriter, r *http.Request) {
		respondWithError(w, 500, "Internal Server Error")
	})

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Printf("Server listening on port %s", port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

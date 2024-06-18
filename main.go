package main

// Importing the SQL driver github.com/lib/pq
// Underscore tells Go that we're importing it for its side effects

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/ggrangel/blog-aggregator/internal/database"
)

type apiConfig struct {
	Db *database.Queries
}

func main() {

	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT must be set")
	}
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Fatal("DATABASE_URL must be set")
	}

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)
	apiConfig := apiConfig{Db: dbQueries}

	go startScraping(dbQueries, 10, time.Minute)

	router := chi.NewRouter()
	v1Router := chi.NewRouter()
	router.Mount("/v1", v1Router)

	v1Router.HandleFunc("GET /healthz", handlerReadiness)
	v1Router.HandleFunc("GET /err", func(w http.ResponseWriter, r *http.Request) {
		respondWithError(w, 500, "Internal Server Error")
	})

	v1Router.HandleFunc("GET /users", apiConfig.middlewareAuth(apiConfig.handlerUsersGet))
	v1Router.HandleFunc("POST /users", apiConfig.handlerUserCreate)

	v1Router.HandleFunc("POST /feeds", apiConfig.middlewareAuth(apiConfig.handlerFeedsCreate))
	v1Router.HandleFunc("GET /feeds", apiConfig.handlerFeedsGet)

	v1Router.HandleFunc(
		"POST /feed_follows",
		apiConfig.middlewareAuth(apiConfig.handlerFeedFollowsCreate),
	)
	v1Router.HandleFunc(
		"DELETE /feed_follows/{id}",
		apiConfig.handlerDeleteFeedFollows)
	v1Router.HandleFunc(
		"GET /feed_follows",
		apiConfig.middlewareAuth(apiConfig.handlerGetFeedFollowsForUser),
	)
	v1Router.HandleFunc(
		"GET /posts", apiConfig.middlewareAuth(apiConfig.handlerGetPostsForUser),
	)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Printf("Server listening on port %s", port)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

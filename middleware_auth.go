package main

import (
	"net/http"

	"github.com/ggrangel/blog-aggregator/internal/auth"
	"github.com/ggrangel/blog-aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Couldn't find API key")
		}

		user, err := cfg.Db.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Coulnd't find user")
			return
		}

		handler(w, r, user)
	}
}

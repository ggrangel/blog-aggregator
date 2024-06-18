package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/ggrangel/blog-aggregator/internal/database"
)

func (apiConfig *apiConfig) handlerUserCreate(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respondWithError(w, http.StatusBadRequest, "Bad request")
	}

	currentTimedate := sql.NullTime{Time: time.Now().UTC(), Valid: true}

	createUserParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: currentTimedate,
		UpdatedAt: currentTimedate,
		Name:      request.Name,
	}

	apiConfig.Db.CreateUser(r.Context(), createUserParams)

	respondWithJson(w, http.StatusOK, createUserParams)
}

func (apiConfig *apiConfig) handlerUserGetByApiKey(w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Authorization")
	apiKey := strings.TrimPrefix(header, "ApiKey ")

	user, err := apiConfig.Db.GetUserByApiKey(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "User not found")
		return
	}

	respondWithJson(w, http.StatusOK, user)
}

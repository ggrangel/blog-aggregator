package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
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

func (apiConfig *apiConfig) handlerUsersGet(
	w http.ResponseWriter,
	r *http.Request,
	user User,
) {
	respondWithJson(w, http.StatusOK, user)
}

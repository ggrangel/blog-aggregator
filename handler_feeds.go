package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/ggrangel/blog-aggregator/internal/database"
)

func (cfg *apiConfig) handlerFeedsCreate(
	w http.ResponseWriter,
	r *http.Request,
	user database.User,
) {
	var request struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respondWithError(w, http.StatusBadRequest, "Bad request")
	}

	userD := databaseUserToUser(user)

	currentTimedate := sql.NullTime{Time: time.Now().UTC(), Valid: true}
	userId := uuid.NullUUID{
		UUID:  userD.ID,
		Valid: true,
	}

	createFeedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: currentTimedate,
		UpdatedAt: currentTimedate,
		UserID:    userId,
		Name:      request.Name,
		Url:       request.Url,
	}

	feed, err := cfg.Db.CreateFeed(r.Context(), createFeedParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed")
		return
	}

	respondWithJson(w, http.StatusOK, databaseFeedToFeed(feed))
}

func (cfg *apiConfig) handlerFeedsGet(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.Db.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get feeds")
		return
	}

	var response []Feed
	for _, feed := range feeds {
		response = append(response, databaseFeedToFeed(feed))
	}

	respondWithJson(w, http.StatusOK, response)
}

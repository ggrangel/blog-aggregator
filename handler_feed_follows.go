package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/ggrangel/blog-aggregator/internal/database"
)

func (cfg *apiConfig) handlerFeedFollowsCreate(w http.ResponseWriter, r *http.Request, user User) {
	var request struct {
		FeedId uuid.NullUUID `json:"feed_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respondWithError(w, http.StatusBadRequest, "Bad request")
	}

	currentTimedate := sql.NullTime{Time: time.Now().UTC(), Valid: true}
	userId := uuid.NullUUID{
		UUID:  user.ID,
		Valid: true,
	}

	feedFollowsParams := database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		UserID:    userId,
		FeedID:    request.FeedId,
		CreatedAt: currentTimedate,
		UpdatedAt: currentTimedate,
	}

	follow, err := cfg.Db.CreateFeedFollows(r.Context(), feedFollowsParams)
	fmt.Println(follow)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusCreated, databaseFeedFollowToFeedFollow(follow))
}

func (cfg *apiConfig) handlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request) {
	followIdString := r.PathValue("id")
	followId, err := uuid.Parse(followIdString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = cfg.Db.DeleteFeedFollows(r.Context(), followId)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal error")
		return
	}

	respondWithJson(w, http.StatusNoContent, nil)
}

func (cfg *apiConfig) handlerGetFeedFollowsForUser(
	w http.ResponseWriter,
	r *http.Request,
	user User,
) {
	follows, err := cfg.Db.GetFeedFollowsForUser(
		r.Context(),
		uuid.NullUUID{UUID: user.ID, Valid: true},
	)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal error")
		return
	}

	feedFollows := make([]FeedFollow, len(follows))
	for i, follow := range follows {
		feedFollows[i] = databaseFeedFollowToFeedFollow(follow)
	}

	respondWithJson(w, http.StatusOK, feedFollows)
}

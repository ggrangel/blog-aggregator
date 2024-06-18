-- name: CreateFeedFollows :one
INSERT INTO follows (id, created_at, updated_at, user_id, feed_id) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: DeleteFeedFollows :exec
DELETE FROM follows WHERE id = $1;

-- name: GetFeedFollowsForUser :many
SELECT * FROM follows WHERE user_id = $1;


-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, description, feed_id, url, published_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetPostsForUser :many
SELECT posts.* from posts
JOIN follows ON posts.feed_id = follows.feed_id
WHERE follows.user_id = $1
ORDER BY posts.published_at DESC
LIMIT $2;

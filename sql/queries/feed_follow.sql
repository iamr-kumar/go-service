-- name: CreateFeedFollow :one
INSERT INTO feed_follows(id, user_id, feed_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;


-- name: GetFeedFollows :many
SELECT * FROM feed_follows WHERE user_id = $1;

-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds 
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT $1;

-- name: MarkFeedFetched :one
UPDATE feeds 
SET last_fetched_at = NOW(), 
updated_at = NOW() WHERE id = $1 
RETURNING *;
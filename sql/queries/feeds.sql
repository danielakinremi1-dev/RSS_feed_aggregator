-- name: CreateFeed :one

INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeedID :one

SELECT id
FROM feeds
WHERE url = $1 LIMIT 1;

-- name: GetFeeds :many

SELECT feeds.name, feeds.url, 
users.name as creator
FROM feeds
INNER JOIN users on users.id = feeds.user_id;

-- name: MarkFeedFetched :exec

UPDATE feeds
SET updated_at = $1, last_fetched_at = $2
WHERE id = $3;

-- name: GetNextFeedToFetch :one

SELECT *
FROM feeds 
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;
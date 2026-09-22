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

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetUserNameByFeedUserID :one
SELECT users.name
FROM users
JOIN feeds ON feeds.user_id = users.id
WHERE feeds.user_id = $1;

-- name: GetFeedByURL :one
SELECT * FROM feeds WHERE url = $1;
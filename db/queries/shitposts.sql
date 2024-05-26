-- name: ListShitposts :many
SELECT id, title, author, content, created_at, updated_at FROM shitposts;

-- name: GetShitpostById :one
SELECT id, title, author, content, created_at, updated_at FROM shitposts WHERE id = ?;

-- name: CreateShitpost :one
INSERT INTO shitposts (title, author, content, passcode) VALUES (?, ?, ?, ?) RETURNING *;

-- name: GetShitpostByIdAndPasscode :one
SELECT id FROM shitposts WHERE id = ? and passcode = ?;

-- name: DeleteShitpostById :exec
DELETE FROM shitposts WHERE id = ? and passcode = ?;

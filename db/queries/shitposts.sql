-- name: ListShitposts :many

SELECT id, author, content, created_at, updated_at FROM shitposts;

-- name: GetShitpostById :one

SELECT id, author, content, created_at, updated_at FROM shitposts WHERE id = ?;

-- name: CreateShitpost :one

INSERT INTO shitposts (author, content, passcode) VALUES (?, ?, ?) RETURNING *;

-- name: DeleteShitpostById :one

DELETE FROM shitposts WHERE id = ? and passcode = ? RETURNING *;

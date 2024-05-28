// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: shitposts.sql

package db

import (
	"context"
	"time"
)

const countShitposts = `-- name: CountShitposts :one
SELECT COUNT(*) FROM shitposts
`

func (q *Queries) CountShitposts(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countShitposts)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createShitpost = `-- name: CreateShitpost :one
INSERT INTO shitposts (title, author, content, passcode) VALUES (?, ?, ?, ?) RETURNING id, title, author, content, passcode, created_at, updated_at
`

type CreateShitpostParams struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Content  string `json:"content"`
	Passcode string `json:"passcode"`
}

func (q *Queries) CreateShitpost(ctx context.Context, arg CreateShitpostParams) (Shitpost, error) {
	row := q.db.QueryRowContext(ctx, createShitpost,
		arg.Title,
		arg.Author,
		arg.Content,
		arg.Passcode,
	)
	var i Shitpost
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Author,
		&i.Content,
		&i.Passcode,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteShitpostById = `-- name: DeleteShitpostById :exec
DELETE FROM shitposts WHERE id = ? and passcode = ?
`

type DeleteShitpostByIdParams struct {
	ID       int64  `json:"id"`
	Passcode string `json:"passcode"`
}

func (q *Queries) DeleteShitpostById(ctx context.Context, arg DeleteShitpostByIdParams) error {
	_, err := q.db.ExecContext(ctx, deleteShitpostById, arg.ID, arg.Passcode)
	return err
}

const getShitpostById = `-- name: GetShitpostById :one
SELECT id, title, author, content, created_at, updated_at FROM shitposts WHERE id = ?
`

type GetShitpostByIdRow struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (q *Queries) GetShitpostById(ctx context.Context, id int64) (GetShitpostByIdRow, error) {
	row := q.db.QueryRowContext(ctx, getShitpostById, id)
	var i GetShitpostByIdRow
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Author,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getShitpostByIdAndPasscode = `-- name: GetShitpostByIdAndPasscode :one
SELECT id FROM shitposts WHERE id = ? and passcode = ?
`

type GetShitpostByIdAndPasscodeParams struct {
	ID       int64  `json:"id"`
	Passcode string `json:"passcode"`
}

func (q *Queries) GetShitpostByIdAndPasscode(ctx context.Context, arg GetShitpostByIdAndPasscodeParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, getShitpostByIdAndPasscode, arg.ID, arg.Passcode)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const listShitposts = `-- name: ListShitposts :many
SELECT id, title, author, content, created_at, updated_at FROM shitposts ORDER BY created_at DESC LIMIT ? OFFSET ?
`

type ListShitpostsParams struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

type ListShitpostsRow struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (q *Queries) ListShitposts(ctx context.Context, arg ListShitpostsParams) ([]ListShitpostsRow, error) {
	rows, err := q.db.QueryContext(ctx, listShitposts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListShitpostsRow
	for rows.Next() {
		var i ListShitpostsRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Author,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

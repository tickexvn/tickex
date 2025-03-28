// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: authors.query.sql

package authors

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const create = `-- name: Create :one
INSERT INTO authors (
  name, bio
) VALUES (
  $1, $2
)
RETURNING id, name, bio
`

type CreateParams struct {
	Name string      `json:"name"`
	Bio  pgtype.Text `json:"bio"`
}

func (q *Queries) Create(ctx context.Context, arg CreateParams) (Author, error) {
	row := q.db.QueryRow(ctx, create, arg.Name, arg.Bio)
	var i Author
	err := row.Scan(&i.ID, &i.Name, &i.Bio)
	return i, err
}

const delete = `-- name: Delete :exec
DELETE FROM authors
WHERE id = $1
`

func (q *Queries) Delete(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, delete, id)
	return err
}

const get = `-- name: Get :one
SELECT id, name, bio FROM authors
WHERE id = $1 LIMIT 1
`

func (q *Queries) Get(ctx context.Context, id int64) (Author, error) {
	row := q.db.QueryRow(ctx, get, id)
	var i Author
	err := row.Scan(&i.ID, &i.Name, &i.Bio)
	return i, err
}

const getMany = `-- name: GetMany :many
SELECT id, name, bio FROM authors
ORDER BY name
`

func (q *Queries) GetMany(ctx context.Context) ([]Author, error) {
	rows, err := q.db.Query(ctx, getMany)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(&i.ID, &i.Name, &i.Bio); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const update = `-- name: Update :one
UPDATE authors
  set name = $2,
  bio = $3
WHERE id = $1
RETURNING id, name, bio
`

type UpdateParams struct {
	ID   int64       `json:"id"`
	Name string      `json:"name"`
	Bio  pgtype.Text `json:"bio"`
}

func (q *Queries) Update(ctx context.Context, arg UpdateParams) (Author, error) {
	row := q.db.QueryRow(ctx, update, arg.ID, arg.Name, arg.Bio)
	var i Author
	err := row.Scan(&i.ID, &i.Name, &i.Bio)
	return i, err
}

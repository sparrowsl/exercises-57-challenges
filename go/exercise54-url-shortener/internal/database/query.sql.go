// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package database

import (
	"context"
)

const allURLs = `-- name: AllURLs :many
SELECT id, short_url, long_url FROM urls
ORDER BY id
`

func (q *Queries) AllURLs(ctx context.Context) ([]Url, error) {
	rows, err := q.db.QueryContext(ctx, allURLs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Url{}
	for rows.Next() {
		var i Url
		if err := rows.Scan(&i.ID, &i.ShortUrl, &i.LongUrl); err != nil {
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

const createURL = `-- name: CreateURL :one
INSERT INTO urls (short_url, long_url)
VALUES (?, ?)
RETURNING id, short_url, long_url
`

type CreateURLParams struct {
	ShortUrl string `db:"short_url" json:"short_url"`
	LongUrl  string `db:"long_url" json:"long_url"`
}

func (q *Queries) CreateURL(ctx context.Context, arg CreateURLParams) (Url, error) {
	row := q.db.QueryRowContext(ctx, createURL, arg.ShortUrl, arg.LongUrl)
	var i Url
	err := row.Scan(&i.ID, &i.ShortUrl, &i.LongUrl)
	return i, err
}

const deleteURL = `-- name: DeleteURL :exec
DELETE FROM urls
WHERE id = ?
`

func (q *Queries) DeleteURL(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteURL, id)
	return err
}

const getOneURL = `-- name: GetOneURL :one
SELECT id, short_url, long_url FROM urls
WHERE short_url = ?
LIMIT 1
`

func (q *Queries) GetOneURL(ctx context.Context, shortUrl string) (Url, error) {
	row := q.db.QueryRowContext(ctx, getOneURL, shortUrl)
	var i Url
	err := row.Scan(&i.ID, &i.ShortUrl, &i.LongUrl)
	return i, err
}

const updateURL = `-- name: UpdateURL :exec
UPDATE urls
SET short_url = ?, long_url = ?
WHERE id = ?
RETURNING id, short_url, long_url
`

type UpdateURLParams struct {
	ShortUrl string `db:"short_url" json:"short_url"`
	LongUrl  string `db:"long_url" json:"long_url"`
	ID       int64  `db:"id" json:"id"`
}

func (q *Queries) UpdateURL(ctx context.Context, arg UpdateURLParams) error {
	_, err := q.db.ExecContext(ctx, updateURL, arg.ShortUrl, arg.LongUrl, arg.ID)
	return err
}

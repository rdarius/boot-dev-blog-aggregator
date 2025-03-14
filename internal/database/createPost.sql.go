// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: createPost.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, url, title, description, feed_id, published_at)
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8
       ) ON CONFLICT (url) DO UPDATE
         SET
             id = $1,
             created_at = $2,
             updated_at = $3,
             title = $5,
             description = $6,
             feed_id = $7,
             published_at = $8
         WHERE posts.url = $4
RETURNING id, created_at, updated_at, title, url, description, published_at, feed_id
`

type CreatePostParams struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Url         string
	Title       string
	Description string
	FeedID      uuid.UUID
	PublishedAt time.Time
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Url,
		arg.Title,
		arg.Description,
		arg.FeedID,
		arg.PublishedAt,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Title,
		&i.Url,
		&i.Description,
		&i.PublishedAt,
		&i.FeedID,
	)
	return i, err
}

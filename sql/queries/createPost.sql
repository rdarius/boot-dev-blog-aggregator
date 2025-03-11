-- name: CreatePost :one
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
RETURNING *;
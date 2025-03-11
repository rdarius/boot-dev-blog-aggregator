-- name: ListFeeds :many
SELECT name, url, (SELECT name FROM users WHERE users.id = feeds.user_id) as user_name FROM feeds;
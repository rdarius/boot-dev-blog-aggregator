// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: feleteFeedFollow.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const deleteFeedFollow = `-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows WHERE feed_follows.user_id  = $1 AND feed_follows.feed_id = (SELECT id FROM feeds WHERE url = $2)
`

type DeleteFeedFollowParams struct {
	UserID uuid.UUID
	Url    string
}

func (q *Queries) DeleteFeedFollow(ctx context.Context, arg DeleteFeedFollowParams) error {
	_, err := q.db.ExecContext(ctx, deleteFeedFollow, arg.UserID, arg.Url)
	return err
}

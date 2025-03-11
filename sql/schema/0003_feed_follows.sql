-- +goose Up
CREATE TABLE feed_follows(
        id UUID PRIMARY KEY,
        created_at TIMESTAMP not null DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP not null,
        user_id UUID not null REFERENCES users(id) ON DELETE CASCADE,
        feed_id UUID not null REFERENCES feeds(id) ON DELETE CASCADE,
        CONSTRAINT UC_User_Feed UNIQUE(user_id, feed_id)
);

-- +goose Down
DROP TABLE feeds;
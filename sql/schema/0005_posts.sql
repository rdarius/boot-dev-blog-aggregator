-- +goose Up
CREATE TABLE posts(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP not null DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP not null,
    title VARCHAR(128) not null,
    url VARCHAR(512) not null UNIQUE,
    description TEXT not null,
    published_at TIMESTAMP not null ,
    feed_id UUID not null REFERENCES feeds(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;
-- +goose Up
CREATE TABLE users(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP not null DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP not null,
    name VARCHAR(64) not null UNIQUE
);

-- +goose Down
DROP TABLE users;
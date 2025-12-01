-- +goose Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    token VARCHAR(36) UNIQUE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ DEFAULT NULL
);

-- +goose Down
DROP TABLE users;
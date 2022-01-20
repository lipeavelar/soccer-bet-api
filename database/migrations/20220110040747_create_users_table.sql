
-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id                VARCHAR(100) NOT NULL UNIQUE,
    name              VARCHAR(120) NOT NULL,
    email             VARCHAR(200) NOT NULL UNIQUE,
    password          VARCHAR(1000) NOT NULL,
    is_admin          BOOLEAN NOT NULL DEFAULT FALSE,
    change_password   BOOLEAN NOT NULL DEFAULT TRUE,
    created_at        TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE IF EXISTS users;
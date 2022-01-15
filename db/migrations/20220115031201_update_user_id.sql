-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE users DROP CONSTRAINT users_pkey;
ALTER TABLE users DROP COLUMN id;
ALTER TABLE users ADD COLUMN id serial;
ALTER TABLE users ADD PRIMARY KEY (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE users DROP CONSTRAINT users_pkey;
ALTER TABLE users DROP COLUMN id;
ALTER TABLE users ADD COLUMN id VARCHAR(100);
ALTER TABLE users ADD PRIMARY KEY (id);
-- +goose StatementEnd

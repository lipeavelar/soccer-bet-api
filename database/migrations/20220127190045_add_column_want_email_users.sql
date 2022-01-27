-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD COLUMN want_email BOOLEAN NOT NULL DEFAULT TRUE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN want_email;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
ALTER TABLE teams ADD COLUMN team_acronym VARCHAR(5) NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE teams DROP COLUMN team_acronym;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS matches (
    id                SERIAL,
    api_match_id      INTEGER NOT NULL,
    home_team         VARCHAR(200) NOT NULL,
    away_team         VARCHAR(200) NOT NULL,
    match_date        DATE NOT NULL,
    season            INTEGER NOT NULL,
    match_day         INTEGER NOT NULL,
    home_team_score   INTEGER,
    away_team_score   INTEGER,
    created_at        TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS matches;
-- +goose StatementEnd

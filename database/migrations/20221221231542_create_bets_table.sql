-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS bets (
    id                SERIAL,
    match_id          INTEGER NOT NULL,
    user_id           INTEGER NOT NULL,
    home_team_score   INTEGER,
    away_team_score   INTEGER,
    created_at        TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS bets;
-- +goose StatementEnd

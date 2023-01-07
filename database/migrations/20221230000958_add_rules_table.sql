-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS rules (
    id                SERIAL,
    rule              VARCHAR(200) NOT NULL,
    code              INTEGER NOT NULL,
    points            INTEGER NOT NULL,
    created_at        TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    PRIMARY KEY(id)
);

INSERT INTO rules(rule, code, points)
  VALUES('HOME_TEAM_SCORE', 1, 3);
INSERT INTO rules(rule, code, points)
  VALUES('AWAY_TEAM_SCORE', 2, 3);
INSERT INTO rules(rule, code, points)
  VALUES('GAME_RESULT', 4, 3);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS rules;
-- +goose StatementEnd

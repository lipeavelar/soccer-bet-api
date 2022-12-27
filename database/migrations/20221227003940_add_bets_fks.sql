-- +goose Up
-- +goose StatementBegin
ALTER TABLE bets
    ADD CONSTRAINT fk_bets_matches FOREIGN KEY (match_id) REFERENCES matches (id);
ALTER TABLE bets
    ADD CONSTRAINT fk_bets_users FOREIGN KEY (user_id) REFERENCES users (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE bets
    DROP CONSTRAINT fk_bets_matches;
ALTER TABLE bets
    DROP CONSTRAINT fk_bets_users;
-- +goose StatementEnd

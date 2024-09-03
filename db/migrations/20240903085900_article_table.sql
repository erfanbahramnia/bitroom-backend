-- +goose Up
-- +goose StatementBegin

ALTER TABLE articles
    ADD COLUMN likes BIGINT[] DEFAULT '{}',
    ADD COLUMN dislikes BIGINT[] DEFAULT '{}',
    ADD COLUMN views BIGINT DEFAULT 0;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE articles
    DROP COLUMN likes,
    DROP COLUMN dislikes,
    DROP COLUMN views;
-- +goose StatementEnd

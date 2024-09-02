-- +goose Up
-- +goose StatementBegin

CREATE TABLE article_properties (
    id SERIAL PRIMARY KEY,
    description TEXT NOT NULL,
    image VARCHAR(225) DEFAULT '',
    article_id INT NOT NULL,
    FOREIGN KEY (article_id) REFERENCES articles(id) ON UPDATE CASCADE ON DELETE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE article_properties;

-- +goose StatementEnd
